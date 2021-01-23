<template lang='pug'>
  div.content
    Modal
    div(class='ui two column divided grid')
      div.row
        div.column(style='transition-delay: 2s;')
          form#ask-form(class='ui form' @submit.prevent="onSubmit" action="/" method="POST")
            h5(class='ui dividing header') お問い合わせ
            div.field.thin
              label Name
              div(class='four fields')
                div.field
                  input(type='text' name='shipping[name]' placeholder='Last Name' maxlength=30 v-model.lazy='lastName')
                  p(v-show='!validateLastName') ※必須です
                div.field
                  input(type='text' name='shipping[name]' placeholder='First Name' maxlength=30 v-model.lazy='firstName')
                  p(v-show='!validateFirstName') ※必須です
            div.thin(class='four fields')
              div.field
                label Type
                select(class='ui fluid dropdown' v-model='selectedType')
                  option(value='' disabled) 個人or会社
                  option(value='private') private
                  option(value='companies') companies
                p(v-show='!validateType') ※必須です
            div.thin(class='four fields')
              div.field
                label State
                select(class='ui fluid dropdown' v-model='selectedState')
                  option(value='' disabled) 地域
                  option(v-for='state in states', :key='state', :value='state') {{ state }}
                p(v-show='!validateState') ※必須です
            div(class='two fields')
              div.field
                label Email
                input(type='text' placeholder='Email' maxlength=200 v-model.lazy='email') 
                p(v-show='!validateEmail') ※正しいメールアドレス表記の必要があります
            div.thin(class='three fields')
              div.field
                label Phone
                input(type='text' placeholder='Phone' maxlength=16 v-model.lazy='phone')
            div.field
              label Message
              textarea(rows='17' maxlength=5000 v-model='message')
              p(v-show='!validateMessage') ※必須です
            button(class='ui inverted orange basic button' type='submit') Submit
        div.column(style='transition-delay: 2.2s;')
          div#office-info
            h5 Our Office
            table
              tr
                td.title
                  p name
                td.value
                  p (株) Example Corporation
              tr
                td.title
                  p founded
                td.value
                  p Since 2019/08/31
              tr
                td.title
                  p ceo
                td.value
                  p Wataru Morioka
              tr
                td.title
                  p phone
                td.value
                  p 03-0000-0000
              tr
                td.title
                  p e-mail
                td.value
                  p example@gmail.com
              tr
                td.title
                  p address
                td.value
                  p 
                    span 〒111-1111
                    br
                    span 東京都渋谷区恵比寿新天地 暫定ビルディング 15F
            br
            div.icon-wrap
              img.icon(src='../assets/icn_fb.png')
              img.icon(src='../assets/icn_insta.png')
              img.icon(src='../assets/icn_tw.png')
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import jQuery from 'jQuery';
import axios from 'axios';
// tslint:disable-next-line:no-var-requires
const https = require('https');
import Modal from '@/components/Modal.vue';
import { setTimeout } from 'timers';
import { states } from '../util/utils';

@Component({
  components: {
    Modal,
  },
})
export default class ContactContent extends Vue {
  private lastName: string = '';
  private lastNameValid: boolean = false;
  private firstName: string = '';
  private firstNameValid: boolean = false;
  private email: string = '';
  private emailValid: boolean = false;
  private phone: string = '';
  private message: string = '';
  private messageValid: boolean = false;
  private selectedType: string = '';
  private typeValid: boolean = false;
  private selectedState: string = '';
  private stateValid: boolean = false;
  private states: string[] = states;

  get validateLastName(): boolean {
    if (this.lastName.length > 0) {
      this.lastNameValid = true;
      return true;
    }
    this.lastNameValid = false;
    return false;
  }

  get validateFirstName(): boolean {
    if (this.firstName.length > 0) {
      this.firstNameValid = true;
      return true;
    }
    this.firstNameValid = false;
    return false;
  }

  get validateType(): boolean {
    if (this.selectedType.length > 0) {
      this.typeValid = true;
      return true;
    }
    this.typeValid = false;
    return false;
  }

  get validateState(): boolean {
    if (this.selectedState.length > 0) {
      this.stateValid = true;
      return true;
    }
    this.stateValid = false;
    return false;
  }

  get validateEmail(): boolean {
    const regexp = /^[A-Za-z0-9]{1}[A-Za-z0-9_.-]*@{1}[A-Za-z0-9_.-]{1,}\.[A-Za-z0-9]{1,}$/;

    if (regexp.test(this.email)) {
      this.emailValid = true;
      return true;
    }
    this.emailValid = false;
    return false;
  }

  get validateMessage(): boolean {
    if (this.message.length > 0) {
      this.messageValid = true;
      return true;
    }
    this.messageValid = false;
    return false;
  }

  private mounted() {
    this.fadein();
    $('.content').scroll(() => {
      this.fadein();
    });
  }

  private fadein(): void {
    const offset = - 60;
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

  private async onSubmit(): Promise<void> {
    if (!this.lastNameValid || !this.firstNameValid || !this.emailValid
     || !this.messageValid || !this.typeValid || !this.stateValid) {
      alert('無効な項目があります');
      return;
    }

    await ($('#loading-modal') as any).modal({
      closable: false,
    }).modal('show');

    const body = {
      name: this.lastName + ' ' + this.firstName,
      organization: this.selectedType,
      state: this.selectedState,
      email: this.email,
      phone: this.phone,
      message: this.message.replace(/^\s+/, '').replace(/\s+$/, ''),
    };

    await axios.post('https://flask.site:443/contact', body, {
      headers: this.$store.state.authHeader,
    })
    .then((res) => {
      if (!res.data.result) {
        alert('送信に失敗しました');
        return;
      }
      this.lastName = '';
      this.firstName = '';
      this.selectedType = '';
      this.selectedState = '';
      this.email = '';
      this.phone = '';
      this.message = '';
      alert('送信が完了しました');
    })
    .catch((err) => {
      console.log(err);
      alert('送信に失敗しました');
    });
    ($('.modal') as any).modal('hide');
  }
}
</script>

<style scoped lang='scss'>
.content {
  position: fixed;
  overflow-y: scroll;
  -webkit-overflow-scrolling: touch;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin: auto;
  height: auto;
  top: 60px; right: 0; bottom: 0; left: 0;
}

::-webkit-scrollbar {
  display: none;
  -webkit-appearance: none;
}

.grid {
  width: 100%;
}

.column {
  min-width: 50%;
  opacity: 0;
  transform: translate(0px, 60px) translate3d(0, 0, 0);
  transition: 1s;
  -ms-filter: blur(50px);
  filter: blur(50px);
}

input, select {
  height: 30px !important; 
}

.button {
  display: block;
  text-align: left;
}

#office-info {
  display: inline-block;

  tr {
    height: 30px;
    td {
      text-align: left;
      font-size: 12px;
    }
  }

  .value {
    padding-left: 50px;
    color: #ffffff;
  }

  .icon-wrap {
    text-align: left;
    .icon {
      cursor: pointer;
      height: 12px;
      margin-left: 10px;
      margin-right: 10px;
    }
  }
}

#ask-form {
  width: 80%;
  margin: auto;

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
}

@media screen and (max-width: 768px){
  .content {
    justify-content: flex-start;
  }

  .column {
    text-align: center;
    min-width: 100%;
    padding-left: 0px;
  }

  table {
    margin: auto;
  }

  #ask-form {
    width: 100%;
    margin-left: 0px;
    margin-bottom: 50px;

    .thin {
      width: 50%;
    }

    input, select {
      font-size: 12px;
      height: 25px !important;
    }

    .button, textarea {
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

  #office-info {
    margin-bottom: 50px;
    tr td {
      font-size: 9px;
    }
  }
}
</style>
