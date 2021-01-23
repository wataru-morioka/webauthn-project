<template lang='pug'>
  div#article-content.contents
    Modal
    ConfirmModal(:confirmMessage='confirmMessage')
    div.article-list(class='ui comments', v-for='([articleId, articleInfo], index) in Array.from(this.$store.state.distinctArticleMap)', :key='articleId')
      input(type='hidden', :value='setArticleId(articleId)')
      div.comment
        a.avatar
          img.avatar(:src='setThumbnail(articleId)')
        div.content
          a.author {{ articleInfo.contributorName }}
          div.metadata
            span.date {{ articleInfo.createdDatetime }}
          div.article.history
            div.button-wrap
            div.toolbar-container
            div.editor
              span(v-html='setBody(articleInfo.body)')
          hr.comment-border
          div.input-comment-wrap
            div.comment-icon(v-if='isVip', @click='showInputComment($event)')
              i(class='pencil alternate icon') 
                span comment
            form(class='ui reply form')
              div.field
                textarea.comment-text
              button(class='ui inverted green button', type='button', @click='sendComment($event, articleId)') Send a Comment
            div(class='ui comments')
              div.comment.history(v-for='(comment, index) in commentArray(articleId)', :key='index')
                a.avatar
                  img.avatar(:src='setCommentThumbnail(comment.thumbnail)')
                div.content
                  a.author {{ comment.name }}
                  div.metadata
                    span.date {{ comment.createdDatetime }}
                  p {{ comment.body }}
            div(style='height: 150px;')
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import { mapState } from 'vuex';
import Modal from '@/components/Modal.vue';
import ConfirmModal from '@/components/ConfirmModal.vue';
import jQuery from 'jQuery';
import axios from 'axios';
import { Comment } from '../model/models';
// tslint:disable-next-line:no-var-requires
const DecoupledEditor = require('@ckeditor/ckeditor5-build-decoupled-document');

@Component({
  components: {
    Modal,
    ConfirmModal,
  },
  computed: mapState({
    isVip: 'isVip',
  }),
})
export default class ArticleContent extends Vue {
  private confirmMessage: string = '';
  private inputPostDisplay: boolean = false;
  private editors: Map<string, any> = new Map<string, any>();
  private articleArray = [];
  private distinctArticleMap = null;
  private isEditing: boolean = false;

  private fadein(): void {
    const offset = 60;
    const effectPos = $(window).height()! - offset;
    $('.article-list').each( function() {
      const thisPos = $(this).offset()!.top;
      if ( effectPos > thisPos ) {
        $(this).css({
            'opacity': 1,
            'transform': 'translate(0px, 0px)',
            '-ms-filter': 'blur(0px)',
            'filter': 'blur(0px)',
        });
      } else {
        $(this).css({
          'opacity': 0,
          'transition-delay': '0s',
        });
      }
    });
  }

  private mounted() {
    this.setEditor();
    setTimeout(() => {
      this.fadein();
      console.log(this.$store.state.distinctArticleMap);
    }, 2500);

    $('.contents').scroll(async () => {
      this.fadein();
      const doch = document.querySelector('#article-content')!.scrollHeight + 120;
      const winh = $(window).innerHeight()!;
      const bottom = doch - winh;

      // スクロールしてページ最下部にきた際、追加でサーバから記事を取得
      if (bottom + 1 === $('#article-content').scrollTop()!) {
        await ($('#loading-modal') as any).modal({
          closable: false,
          onVisible: async () => {
            await this.$store.dispatch('getArticles', {
              additional: true,
            });

            await ($('.modal') as any).modal({
              closable: false,
            }).modal('hide');
          },
        }).modal('show');
      }
    });
  }

  private setBody(body: string): string {
    return body;
  }

  // コメント入力エリアを表示、非表示
  private showInputComment(event: any): void {
    const parent = $(event.currentTarget).closest('.input-comment-wrap');
    const form = $(parent).children('form')[0];

    if ($(form).css('display') === 'none') {
      $(form).slideToggle(300);
    } else {
      $(form).hide(300);
    }
  }

  // 記事に紐づくコメントリストを取得
  private commentArray(articleId: number): Comment[] {
    const targetArray = this.$store.state.articleArray.filter((article: any) => {
      return article.id === articleId && article.commentatorName != null;
    });

    const commentArray = targetArray.map((article: any) => {
      return new Comment(
        article.commentatorName,
        article.commentBody,
        article.commentatorThumbnail,
        article.commentCreatedDatetime,
      );
    });

    return commentArray;
  }

  // コメント送信
  private async sendComment(event: any, id: number): Promise<void> {
    const form = $(event.currentTarget).parents('form')[0];
    const commnetArea = $(form).find('textarea')[0];
    let comment = $(commnetArea).val() as string;
    // コメント末尾の空白削除
    comment = comment.replace(/^\s+/, '').replace(/\s+$/, '');

    if (comment.length === 0) {
      alert('コメントを入力してください');
      return;
    }

    if (comment.length > 1000) {
      alert('最大文字数は1000文字です');
      return;
    }

    this.confirmMessage = 'will you send a comment really?';
    ($('#confirm-modal') as any).modal({
      closable: false,
      onApprove: async (el: any) => {
        ($('#loading-modal') as any).modal({
          closable: false,
          onVisible: async () => {
            setTimeout(async () => {
              const body = {
                articleId: id,
                body: comment,
              };

              await axios.post('https://django.service/api/service/comment', body, {
                headers: this.$store.state.authHeader,
              })
              .then(async (res) => {
                if (!res.data.result) {
                  alert('コメント送信に失敗しました');
                  return;
                }

                $(commnetArea).val('');
                $(form).hide(300);

                await this.$store.dispatch('getArticles', {
                  additional: false,
                });

                alert('送信しました');
              })
              .catch((err) => {
                console.log(err);
                alert('コメント送信に失敗しました');
              });

              ($('.modal') as any).modal('hide');
            }, 500);
          },
        }).modal('show');
      },
    }).modal('show');
  }

  private setArticleId(id: number): number {
    return id;
  }

  // 記事、コメントが更新された時
  private updated() {
    this.fadein();
    // CKEdirotを再セット
    this.resetEditor();
  }

  // CKEditorセット（readonlyモード）
  private setEditor(): void {
    $('.article.history').each((index, element) => {
      const toolbarElement = $(element).children('.toolbar-container')[0];
      const editorElement = $(element).children('.editor')[0];
      const inputElement = $(element).parents('.article-list').children('input')[0];
      const articleId: string = $(inputElement).val() as string;

      DecoupledEditor
      .create( editorElement, {
      })
      .then( (editor: any) => {
        this.editors.set(articleId, editor);
        editor.isReadOnly = true;

        // 記事内容をセット
        const target: any = this.$store.state.articleArray.filter((x: any) => {
          return (x as any).id === Number(articleId);
        });
        editor.setData(target[0].body);
      })
      .catch( (err: any) => {
        console.error( err.stack );
      });
    });
  }

    // 記事が追加、変更された場合、CKEditorを再セット
  private resetEditor(): void {
    // 過去の記事一覧エリア
    $('.article.history').each((index, element) => {
      const toolbarElement = $(element).children('.toolbar-container')[0];
      const editorElement = $(element).children('.editor')[0];
      const inputElement = $(element).parents('.article-list').children('input')[0];
      const articleId: string = $(inputElement).val() as string;
      const target: any = this.$store.state.articleArray.filter((x: any) => {
        return (x as any).id === Number(articleId);
      });

      if (this.editors.get(articleId)) {
        // 一旦インスタンス破棄
        this.editors.get(articleId).destroy();
      }

      DecoupledEditor
        .create( editorElement, {
          ckfinder: {
            uploadUrl: 'https://django.service/api/service/image',
          },
        })
        .then(async (editor: any) => {
          editor.isReadOnly = true;
          this.editors.set(articleId, editor);
          await editor.setData(target[0].body);
        })
        .catch( (err: any) => {
          console.log(err);
          console.error( err.stack );
        });
    });
  }

  // コメント送信者のサムネイルをセット
  private setCommentThumbnail(thumbnail: Uint8Array): string {
    const buffer = Buffer.from(thumbnail);
    const blob = new Blob([buffer], {type: (thumbnail as any).mimetype});
    const blobURL = window.URL.createObjectURL(blob);
    return blobURL;
  }

  // 記事投稿者のサムネイルをセット
  private setThumbnail(articleId: number = 0): string {
    const target: any = this.$store.state.articleArray.filter((x: any) => {
      return (x as any).id === articleId;
    });

    const thumbnail = target[0].thumbnail;
    const buffer = Buffer.from(thumbnail);
    const blob = new Blob([buffer], {type: thumbnail.mimetype});
    const blobURL = window.URL.createObjectURL(blob);
    return blobURL;
  }
}
</script>

<style scoped lang='scss'>
.contents {
  position: fixed;
  top: 60px; right: 0; bottom: 60px; left: 0;
  overflow-y: scroll;
  -webkit-overflow-scrolling: touch;
  width: 100%;
  height: auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: auto;
}

::-webkit-scrollbar {
  display: none;
  -webkit-appearance: none;
}

.article-list {
  width: 60%;
  max-width: 100% !important;
  opacity: 0;
  transform: translate(0px, 60px) translate3d(0, 0, 0);
  transition: 1.5s;
  -ms-filter: blur(50px);
  filter: blur(50px);

  .avatar {
    width: 30px !important;
    height: 30px !important;
    border-radius: 50% !important;
  }

  .avatar:hover {
    cursor: pointer;
  }

  .content {
    text-align: left;

    .author {
      color: #ffffff !important;
      font-size: 12px;
    }

    .metadata {
      color: #ffffff77 !important;
      margin-left: 20px;
    }

    .article {
      margin-top: 10px;

      #editor, .editor {
        color: #ffffff;
      }

      .button-wrap {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        align-items: flex-end;

        button {
          margin-top: 10px;
          margin-bottom: 10px;
          font-size: 12px;
        }
      }
    }

    .comment-border {
      height: 1px;
      background-color: #ffffff77;
      border: none;
      margin-top: 10px;
      margin-bottom: 5px;
    }

    .input-comment-wrap {
      .comment-icon {
        margin: 10px;
        // text-align: center;
        span {
          margin-left: 10px;
        }
      }

      .comment-icon:hover {
        cursor: pointer;
      }

      form {
        display: none;

        textarea {
          height: 15px;
        }

        button {
          font-size: 10px;
        }
      }
    }

    .comment.history {
      margin-bottom: 20px;
      p {
        white-space:pre-wrap;
        color: #ffffff;
      }
    }
  }

  img {
    max-width: 100% !important;
  }
}

@media screen and (max-width: 768px){
  .article-list {
    width: 98%;

    .avatar {
      width: 30px !important;
      height: 30px !important;
    }
  }
}
</style>
