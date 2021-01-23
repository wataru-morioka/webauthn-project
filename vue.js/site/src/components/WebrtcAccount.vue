<template lang='pug'>
  div.content
    Modal
    ConfirmModal(:confirmMessage='confirmMessage')
    div#account-form
      form(class='ui form' @submit.prevent="onSubmit" action="/" method="POST")
        h5(class='ui dividing header') Your Account
        div.field.middle
          label name
          input.no-change(type='text', :value='this.$store.state.displayName', readonly)
        div.field.middle
          label account
          input.no-change(type='text', :value='this.$store.state.email', readonly)
        div.thin(class='two fields')
          div.field
            label state
            select(class='ui fluid dropdown' v-model='selectedState')
              option(value='' disabled) 地域
              option(v-for='state in states', :key='state', :value='state') {{ state }}
            p(v-show='!validateState') ※必須です
        div.field.thumbnail
          label thumbnail
          button.thumbnail(class='ui inverted yellow button', type='button', @click='selectThumbnail($event)') thumnail
          input#thumbnail-input(type='file', accept='image/png,image/jpeg,image/gif', @change='onChangeThumbnail($event)')
          br
          img#thumbnail(:src='setThumbnail()')
          p(v-show='!thumbnailValid') ※必須です
        button.submit(class='ui inverted green button', type='submit') Register
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import Modal from '@/components/Modal.vue';
import ConfirmModal from '@/components/ConfirmModal.vue';
import jQuery from 'jQuery';
import axios from 'axios';
import { states } from '../util/utils';

@Component({
  components: {
    Modal,
    ConfirmModal,
  },
})
export default class WebrtcAccount extends Vue {
  private confirmMessage: string = '';
  private thumbnailValid: boolean = false;
  private selectedState: string = '';
  private stateValid: boolean = false;
  private thumbnailBlobUrl: string = '';
  private states: string[] = states;

  private mounted() {
    setTimeout(() => {
      this.fadein();
    }, 1);

    if (this.$store.state.thumbnail != null) {
      this.thumbnailValid = true;
    }

    if (this.$store.state.state != null) {
      this.selectedState = this.$store.state.state;
    }
  }

  private fadein(): void {
    $('#account-form').css({
      'opacity': 1,
      'transform': 'translate(0px, 0px)',
      '-ms-filter': 'blur(0px)',
      'filter': 'blur(0px)',
    });
  }

  get validateState(): boolean {
    if (this.selectedState.length > 0) {
      this.stateValid = true;
      return true;
    }
    this.stateValid = false;
    return false;
  }

  private selectThumbnail(event: any): void {
    $('#thumbnail-input').click();
  }

  private onChangeThumbnail(event: any): void {
    this.thumbnailValid = true;
    const blobUrl = window.URL.createObjectURL(event.target.files[0]);
    this.thumbnailBlobUrl = blobUrl;
    (document.getElementById('thumbnail')! as HTMLImageElement).src = blobUrl;
  }

  private setThumbnail(): string {
    const thumbnail = this.$store.state.thumbnail;
    if (thumbnail === null) {
      return '';
    }

    if (this.thumbnailBlobUrl.length !== 0) {
      return this.thumbnailBlobUrl;
    }

    const buffer = Buffer.from(thumbnail);
    const blob = new Blob([buffer], {type: thumbnail.mimetype});
    const blobURL = window.URL.createObjectURL(blob);
    return blobURL;
  }

  private async onSubmit(): Promise<void> {
    if (!this.stateValid || !this.thumbnailValid) {
      alert('無効な項目があります');
      return;
    }

    const file = $('#thumbnail-input').prop('files')[0];
    if (file && file.size > 1000000) {
      alert('最大容量は1MBです');
      return;
    }

    this.confirmMessage = 'will you register account really?';
    ($('#confirm-modal') as any).modal({
      closable: false,
      onApprove: async (el: any) => {
        ($('#loading-modal') as any).modal({
          closable: false,
          onVisible: async () => {
            const body = new FormData();

            if (file) {
              body.append('file', file);
            }
            body.append('state', this.selectedState);

            await axios.put('https://django.service/api/service/registerVipAccount', body, {
              headers: this.$store.state.authHeader,
            })
            .then((res: any) => {
              if (res.data.result) {
                alert('アップロードが完了しました');
                console.log(res.data);

                const thumbnailBase64 = res.data.thumbnail;
                const bin = atob(thumbnailBase64.replace(/^.*,/, ''));
                const buffer = new Uint8Array(bin.length);

                for (let i = 0; i < bin.length; i++) {
                    buffer[i] = bin.charCodeAt(i);
                }

                this.$store.commit('setUserInfo', {
                  isLogin: true,
                  isAnonymous: false,
                  thumbnail: buffer,
                  state: res.data.state,
                  isVip: true,
                  isAdmin: this.$store.state.isAdmin,
                });
              } else {
                alert('photoアップロードに失敗しました');
              }
            })
            .catch((err) => {
              alert(err);
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
  top: 60px; right: 0; bottom: 60px; left: 0;
  width: 100%;
  height: auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin: auto;
}

#account-form {
  width: 300px;
  opacity: 0;
  transform: translate(0px, 60px) translate3d(0, 0, 0);
  transition: 1s;
  -ms-filter: blur(50px);
  filter: blur(50px);

  .field {
    margin-bottom: 30px;
  }

  #thumbnail-input {
    display: none;
  }

  .no-change {
    background: #545454;
    color: #ffffff;
  }

  label {
    text-align: left;
    color: #ffffff77 !important;
    font-size: 11px !important;
  }

  h5 {
    text-align: left;
    color: #ffffff !important;
    font-size: 12px !important;
  }

  p {
    text-align: left;
    color: #b14c2d;
    font-size: 12px;
  }

  .thumbnail {
    p {
      text-align: center;
    }
  }

  img {
    width: 100px;
    height: 100px;
    border-radius: 50%;
  }

  button.thumbnail {
    font-size: 10px;
    margin-bottom: 10px;
  }

  button.submit {
    display: block;
    margin-top: 40px;
    // margin-left: auto;
    margin-right: 0px;
  }
}

@media screen and (max-width: 768px){
  .content {
    bottom: 0px;

    #account-form {
      width: 95%;
      
      .field {
        margin-bottom: 10px;
      }

      .thin {
        width: 40%;
      }

      .middle {
        width: 80%;
      }
      
      .button {
        font-size: 12px;
      }

      h5 {
        font-size: 10px !important;
      }

      p {
        font-size: 10px;
      }

      label {
        font-size: 8px !important;
      }
    }
  }
}
</style>