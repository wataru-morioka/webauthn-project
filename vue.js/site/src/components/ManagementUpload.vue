<template lang='pug'>
div
  Modal
  ConfirmModal(:confirmMessage='confirmMessage')
  div#sub-menu-wrap
    SubMenu
  video#project-video(controls playsinline)
  div#add-image
    button(type=button class='ui inverted red button', @click='addPhoto($event)') add
    input#add-input(type='file', accept='image/png,image/jpeg,image/gif', @change='onChangeAddPhoto($event)')
  div.content(class='ui two column divided grid', @click='stop')
    div.row(v-for='(photoArray, index) in photoMultiArray', :key='index', :style='firstChildRow(index)')
      div.column.left(:style='transitionDelay(0, 0.2, index * 2)')
        div.content-box.box-left
          div.edit-photo-wrap
            button(type=button class='ui inverted primary button', @click='editPhoto($event)') edit
            div.upload-wrap
              button(type=button class='ui inverted orange button', @click='uploadThumbnail($event)') thumbnail
              input.edit-thumbnail(type='file', accept='image/png,image/jpeg,image/gif', @change='onChangeThumbnail($event, photoArray[0].id)')
              button(type=button class='ui inverted orange button', @click='uploadVideo($event)') video
              input.edit-video(type='file', accept='video/mp4', @change='onChangeVideo($event, photoArray[0].id)')
          div.image(@click='play(photoArray[0].id)')
            i.huge.play.icon
            img(:src='setPhoto(index, 0)')
          div.subject
            input.sub-title(type='text', readonly=true, :value='photoArray[0].subTitle')
            button(type=button class='ui inverted primary button', @click='editSubTitle($event, photoArray[0].id)') edit
            h4 
              input.title(type='text', readonly=true, :value='photoArray[0].title')
              button(type=button class='ui inverted primary button', @click='editTitle($event, photoArray[0].id)') edit
            button(type=button class='ui inverted yellow button', @click='minify(photoArray[0].id)') minify
            button(type=button class='ui inverted green button', @click='download(photoArray[0].id)') download
            button(type=button class='ui inverted red button', @click='deletePhoto(photoArray[0].id)') delete
      div.column(v-if='existsPhoto(photoArray[1])', :style='transitionDelay(0, 0.2, index * 2 + 1)')
        div.content-box.box-right
          div.edit-photo-wrap
            button(type=button class='ui inverted primary button', @click='editPhoto($event)') edit
            div.upload-wrap
              button(type=button class='ui inverted orange button', @click='uploadThumbnail($event)') thumbnail
              input.edit-thumbnail(type='file', accept='image/png,image/jpeg,image/gif', @change='onChangeThumbnail($event, photoArray[1].id)')
              button(type=button class='ui inverted orange button', @click='uploadVideo($event)') video
              input.edit-video(type='file', accept='video/mp4', @change='onChangeVideo($event, photoArray[1].id)')
          div.image(@click='play(photoArray[1].id)')
            i.huge.play.icon
            img(:src='setPhoto(index, 1)')
          div.subject
            input.sub-title(type='text', readonly=true, :value='photoArray[1].subTitle')
            button(type=button class='ui inverted primary button', @click='editSubTitle($event, photoArray[1].id)') edit
            h4 
              input.title(type='text', readonly=true, :value='photoArray[1].title')
              button(type=button class='ui inverted primary button', @click='editTitle($event, photoArray[1].id)') edit
            button(type=button class='ui inverted yellow button', @click='minify(photoArray[1].id)') minify
            button(type=button class='ui inverted green button', @click='download(photoArray[1].id)') download
            button(type=button class='ui inverted red button', @click='deletePhoto(photoArray[1].id)') delete
</template>

<script lang='ts'>
import Vue from 'vue';
import { Component, Mixin, Mixins } from 'vue-mixin-decorator';
import { mapState } from 'vuex';
import Modal from '@/components/Modal.vue';
import ConfirmModal from '@/components/ConfirmModal.vue';
import SubMenu from '@/components/SubMenu.vue';
import VideoMixin from '@/components/VideoMixin.vue';
import axios from 'axios';
import jQuery from 'jQuery';
import firebase from 'firebase/app';
import { PhotoInfo } from '../model/models';

@Component({
  components: {
    Modal,
    ConfirmModal,
    SubMenu,
  },
})
export default class ManagementUpload extends Mixins<VideoMixin>(VideoMixin) {
  private confirmMessage: string = '';
  private photoMultiArray: PhotoInfo[][] = [];

  private transitionDelay(diff: number, rate: number, index: number): string {
    const delay = diff + rate * index;
    return 'transition-delay: ' + delay + 's';
  }

  private firstChildRow(index: number): string {
    if (index === 0) {
      return 'margin-top: 70px;';
    }
    return 'margin-top: auto;';
  }

  // プロジェクト動画のサムネイルセット
  private setPhoto(indexX: number, indexY: number): string {
    const buffer = Buffer.from(this.photoMultiArray[indexX][indexY].data);
    const blob = new Blob([buffer], {type: this.photoMultiArray[indexX][indexY].mimetype});
    const blobURL = window.URL.createObjectURL(blob);
    return blobURL;
  }

  private async mounted() {
    this.fadein();
    $('.content').scroll(() => {
      this.fadein();
      (this as any).stop();
    });

    $('.project-row').each((index) => {
      if (index !== 0) { return; }

      $(this).css({
        'padding-top': '70px !important',
      });
    });

    // プロジェクト情報を取得
    await this.$store.dispatch('getPhotos');
    this.photoMultiArray = this.$store.getters.getPhotos;
  }

  private created() {
    this.photoMultiArray = this.$store.getters.getPhotos;
    this.fadein();
    this.$store.commit('setInitVideoFlag');
  }

  private fadein(): void {
    const offset = 30;
    const effectPos = $(window).height()! - offset;
    $('.column').each( function() {
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

  // 対象の現在の動画サムネイルを取得
  private async download(id: number): Promise<void> {
    axios.get('https://express.management/download', {
      headers: this.$store.state.authHeader,
      params: {
        photoId: id,
      },
    })
    .then((res: any) => {
      const photo = res.data.photoInfo;
      const buffer = Buffer.from(photo.data);
      const blob = new Blob([buffer], {type: photo.mimetype});
      const blobURL = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.download = photo.file_name;
      a.href = blobURL;
      a.click();
    })
    .catch((err) => {
      alert(err);
    });
  }

  // サムネイルを圧縮するリクエストをサーバに送信し、完了後プロジェクト情報を更新する
  private async minify(id: number): Promise<void> {
    this.confirmMessage = 'will you minify thumbnail really?';
    ($('#confirm-modal') as any).modal({
      closable: false,
      onApprove: async (el: any) => {
        ($('#loading-modal') as any).modal({
          closable: false,
          onVisible: async () => {
            await axios.get('https://express.management/minify', {
              headers: this.$store.state.authHeader,
              params: {
                photoId: id,
              },
            })
            .then((res: any) => {
              alert('圧縮が完了しました');
            })
            .catch((err) => {
              alert('圧縮に失敗しました');
            });

            await this.$store.dispatch('getPhotos');
            this.photoMultiArray = this.$store.getters.getPhotos;

            ($('.modal') as any).modal('hide');
          },
        }).modal('show');
      },
    }).modal('show');
  }

  // サムネイル、動画のeditボタン押下時
  private editPhoto(event: any): void {
    const target = event.currentTarget;
    const parent = $(target).closest('.edit-photo-wrap');
    const uploadArea = $(parent).children('.upload-wrap');

    // css変更（ボタンエリア表示、非表示）
    if ($(target).hasClass('primary')) {
      (uploadArea).slideToggle(300);
      $(target).removeClass('primary');
      $(target).addClass('secondary');
      $(target).text('cancel');
      return;
    }
    $(uploadArea).hide(300);
    $(target).removeClass('secondary');
    $(target).addClass('primary');
    $(target).text('edit');
  }

  // 変更サムネイル洗選択完了時
  private async onChangeThumbnail(event: any, id: number): Promise<void> {
    const target = event.currentTarget;
    this.confirmMessage = 'will you change photograph really?';

    // 確認モーダル表示
    ($('#confirm-modal') as any).modal({
      closable: false,
      onApprove: async (el: any) => {
        ($('#loading-modal') as any).modal({
          closable: false,
          onVisible: async () => {
            const body = new FormData();
            body.append('file', event.target.files[0]);
            body.append('photoId', String(id));

            await axios.put('https://express.management/photographs', body, {
              headers: this.$store.state.authHeader,
            })
            .then((res: any) => {
              alert('アップロードが完了しました');
            })
            .catch((err) => {
              alert('アップロードに失敗しました');
            });

            const parent = $(target).closest('.edit-photo-wrap');
            const uploadArea = $(parent).children('.upload-wrap');
            const editButton = $(parent).children('button')[0];

            $(uploadArea).hide(300);
            $(editButton).removeClass('secondary');
            $(editButton).addClass('primary');
            $(editButton).text('edit');

            // 画面更新
            await this.$store.dispatch('getPhotos');
            this.photoMultiArray = this.$store.getters.getPhotos;

            ($('.modal') as any).modal('hide');
            $(target).val('');
          },
        }).modal('show');
      },
      onDeny: (el: any) => {
        $(target).val('');
      },
    }).modal('show');
  }

  // 動画選択完了時
  private async onChangeVideo(event: any, id: number): Promise<void> {
    const target = event.currentTarget;
    this.confirmMessage = 'will you change video really?';

    ($('#confirm-modal') as any).modal({
      closable: false,
      onApprove: async (el: any) => {
        ($('#loading-modal') as any).modal({
          closable: false,
          onVisible: async () => {
            const body = new FormData();
            const file = event.target.files[0];

            if (file.size > 100000000) {
              alert('アップロード最大サイズは「100MB」です');
              $(target).val('');
              setTimeout(() => {
                ($('.modal') as any).modal('hide');
              }, 1);
              return;
            }

            body.append('file', event.target.files[0]);
            body.append('photoId', String(id));

            await axios.put('https://express.management/video', body, {
              headers: this.$store.state.authHeader,
            })
            .then((res: any) => {
              alert('動画アップロードに成功しました');
            })
            .catch((err) => {
              alert('動画アップロードが失敗しました');
            });

            const parent = $(target).closest('.edit-photo-wrap');
            const uploadArea = $(parent).children('.upload-wrap');
            const editButton = $(parent).children('button')[0];

            $(uploadArea).hide(300);
            $(editButton).removeClass('secondary');
            $(editButton).addClass('primary');
            $(editButton).text('edit');

            // 画面更新
            await this.$store.dispatch('getPhotos');
            this.photoMultiArray = this.$store.getters.getPhotos;

            ($('.modal') as any).modal('hide');
            $(target).val('');
          },
        }).modal('show');
      },
      onDeny: (el: any) => {
        $(target).val('');
      },
    }).modal('show');
  }

  // サムネイル更新ボタン押下時
  private async uploadThumbnail(event: any, id: number): Promise<void> {
    const target = event.currentTarget;
    const parent = $(target).closest('.upload-wrap');
    const input = $(parent).children('.edit-thumbnail');
    $(input).click();
  }

  // 動画更新ボタン押下時
  private async uploadVideo(event: any, id: number): Promise<void> {
    const target = event.currentTarget;
    const parent = $(target).closest('.upload-wrap');
    const input = $(parent).children('.edit-video');
    $(input).click();
  }

  // サブタイトルedit、confirmボタン押下時
  private async editSubTitle(event: any, id: number): Promise<void> {
    const target = event.currentTarget;
    const parent = $(target).closest('.subject');
    const input = $(parent).children('.sub-title');

    // 編集→確定ボタンに変更
    if ($(target).hasClass('primary')) {
      $(input).prop('readonly', false);
      $(target).removeClass('primary');
      $(target).addClass('orange');
      $(target).text('confirm');
      return;
    }

    // 確定ボタン押下時
    $(target).prop('disabled', true);
    $(target).addClass('loading');

    const body = {
      photoId: id,
      subTitle: $(input).val(),
    };

    // サーバに変更情報送信
    await axios.put('https://express.management/photographs', body, {
      headers: this.$store.state.authHeader,
    })
    .then(async (res: any) => {
      // プロジェクト情報更新
      await this.$store.dispatch('getPhotos');
      this.photoMultiArray = this.$store.getters.getPhotos;

      // 確定→編集ボタンに変更
      $(input).prop('readonly', true);
      $(target).removeClass('orange');
      $(target).addClass('primary');
      $(target).text('edit');

      alert('更新に成功しました');
    })
    .catch((err) => {
      alert('更新に失敗しました');
    });

    $(target).prop('disabled', false);
    $(target).removeClass('loading');
  }

  // タイトルedit、confirmボタン押下時
  private async editTitle(event: any, id: number): Promise<void> {
    const target = event.currentTarget;
    const parent = $(target).closest('h4');
    const input = $(parent).children('.title');

    // 編集→確定ボタンに変更
    if ($(target).hasClass('primary')) {
      $(input).prop('readonly', false);
      $(target).removeClass('primary');
      $(target).addClass('orange');
      $(target).text('confirm');
      return;
    }

    // 確定ボタン押下時
    $(target).prop('disabled', true);
    $(target).addClass('loading');

    const body = {
      photoId: id,
      title: $(input).val(),
    };

    // サーバに変更情報送信
    await axios.put('https://express.management/photographs', body, {
      headers: this.$store.state.authHeader,
    })
    .then(async (res: any) => {
      // プロジェクト情報更新
      await this.$store.dispatch('getPhotos');
      this.photoMultiArray = this.$store.getters.getPhotos;

      // 確定→編集ボタンに変更
      $(input).prop('readonly', true);
      $(target).removeClass('orange');
      $(target).addClass('primary');
      $(target).text('edit');
      alert('更新に成功しました');
    })
    .catch((err) => {
      alert('更新に失敗しました');
    });

    $(target).removeClass('loading');
    $(target).prop('disabled', false);
  }

  private existsPhoto(photo: any): boolean {
    if (photo === undefined) {
      return false;
    }
    return true;
  }

  private addPhoto(event: any): void {
    $('#add-input').click();
  }

  private addVideo(event: any): void {
    $('#add-video-input').click();
  }

  // 新規プロジェクトサムネイル選択時
  private onChangeAddPhoto(event: any): void {
    this.confirmMessage = 'will you add photograph really?';

    ($('#confirm-modal') as any).modal({
      closable: false,
      onApprove: async (el: any) => {
        ($('#loading-modal') as any).modal({
          closable: false,
          onVisible: async () => {
            const body = new FormData();
            body.append('file', event.target.files[0]);

            await axios.post('https://express.management/photographs', body, {
              headers: this.$store.state.authHeader,
            })
            .then(async (res) => {
              await this.$store.dispatch('getPhotos');
              this.photoMultiArray = this.$store.getters.getPhotos;
              alert('アップロードが完了しました');
            })
            .catch((err) => {
              alert('アップロードに失敗しました');
            });

            ($('.modal') as any).modal('hide');
            $('#add-input').val('');
          },
        }).modal('show');
      },
      onDeny: (el: any) => {
        $('#add-input').val('');
      },
    }).modal('show');
  }

  // deleteボタン押下時
  private deletePhoto(id: number) {
    this.confirmMessage = 'will you delete photograph really?';
    ($('#confirm-modal') as any).modal({
      closable: false,
      onApprove: async (el: any) => {
        ($('#loading-modal') as any).modal({
          closable: false,
          onVisible: async () => {
            await axios.delete('https://express.management/photographs', {
              headers: this.$store.state.authHeader,
              params: {
                photoId: id,
              },
            })
            .then(async (res: any) => {
              await this.$store.dispatch('getPhotos');
              this.photoMultiArray = this.$store.getters.getPhotos;
              alert('削除が完了しました');
            })
            .catch((err) => {
              alert('削除に失敗しました');
            });

            ($('.modal') as any).modal('hide');
          },
        }).modal('show');
      },
    }).modal('show');
  }
}
</script>

<style scoped lang='scss'>
.content {
  position: fixed;
  top: 54px; right: 0; bottom: 0; left: 0;
  overflow-y: scroll;
  -webkit-overflow-scrolling: touch;
  height: auto;
  margin: auto;
  transition: 1s;
  will-change: transform;
}

::-webkit-scrollbar {
  display: none;
  -webkit-appearance: none;
}

#add-image {
  position: fixed;
  top: 54px; right: 0; bottom: 0; left: 0;
  margin-top: 0px !important;
  margin: auto;
  width: 200px;
  height: 20px;
  z-index: 10;

  button {
    font-size: 12px;
  }

  input {
    height: 20px;
    display: none;
  }
}

.edit-thumbnail, .edit-video {
  display: none;
}

.edit-photo-wrap {
  flex-direction: column;
  display: flex;
  button {
    margin: auto;
    margin-left: 0px;
    margin-bottom: 2px;
    font-size: 10px !important;
  }
}

.upload-wrap {
  display: none;
  button {
    margin-left: 2px;
    margin-right: 2px;
    margin-bottom: 2px;
    font-size: 10px !important;
  }
}

#sub-menu-wrap {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  flex-direction: column;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: auto;
  width: 100px;
  height: 270px;
  overflow-y: scroll;
  -webkit-overflow-scrolling: touch;
  z-index: 2;
}

video {
  opacity: 0;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  width: 40%;
  transition: 1s;
  z-index: -10;
}

.image {
  cursor: pointer;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  i {
    position: fixed;
  }
}

img {
  width: 100%;
}

.box-left {
  margin-left: auto;
  margin-right: 150px;
}

.box-right {
  margin-top: 70px;
  margin-left: 150px;
  margin-bottom: 30px;
}
  
.column {
  min-width: 50%;
  width: 100%;
  // height: 100%;
  opacity: 0;
  transform: translate(0px, 30px) translate3d(0, 0, 0);
  transition: 1s;
  will-change: transform;
  -ms-filter: blur(30px);
  filter: blur(30px);
  
  .content-box {
    // position: relative;
    width: 340px;
    // z-index: 5;

    .subject {
      cursor: default;
      text-align: left;
      margin-top: 10px;
      input {
        color: #ffffff;
        font-size: 12px;
        margin-bottom: 0px;
        background: (0, 0, 0, 0) !important;
      }

      .primary {
        // font-size: 10px;
        margin-left: 5px;
      }

      h4 {
        margin-top: 0px;
        input {
          color: #ffffff;
          font-size: 16px;
          display: inline-block;
          transition: 0.6s;
          background: (0, 0, 0, 0) !important;
        }

        button {
          // font-size: 10px;
          margin-left: 5px;
        }
      }

      button {
        font-size: 10px;
      }
    }
  }
}

.content-box:hover {
  h4 p {
    -webkit-transform: rotateX(360deg);
    transform: rotateX(360deg);
  }
}

@media screen and (max-width: 768px){
  .content {
    top: 90px;

    .row {
      margin: auto !important;
      padding-top: 0px;
    }
  }

  #sub-menu-wrap {
    display: none;
  }

  video {
    width: 100%;
  }

  .box-left {
    margin: auto;
  }

  .box-right {
    margin: auto;
  }

  .column {
    min-width: 100%;
    margin-bottom: 30px !important;
    .content-box {
      width: 80%;

      .subject {
        margin-top: 0px;

        p {
          font-size: 6px;
        }

        h4 p {
          padding-top: 0px;
          font-size: 13px;
        }
      }
    }
  }

  .column.left {
    margin-bottom: 60px !important;
  }
}
</style>
