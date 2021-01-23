<template lang='pug'>
  div#footer-nav
    div
      div#login-out
        a(href='#', @click.stop.prevent='login', v-show='!isLogin') Login
        div(v-show='!isLogin', style='height: 10px;')
        a(href='#', @click.stop.prevent='logout', v-show='isLogin') Logout
      p#user-account(v-if='isLogin') {{ email }}
      div.horizontal-array
        div
          img.icon(src='../assets/icn_fb.png')
        div
          img.icon(src='../assets/icn_insta.png')
        div
          img.icon(src='../assets/icn_tw.png')
      a(href='https://www.jagermeister.com/en')
        img#footer-icon(src='../assets/jager-logo.png')
      img#footer-icon2(src='../assets/footer-logo.png')
      //- br
      a.extension(v-if='isAdmin', href='#', @click.stop.prevent='toManagement')
        p(style='transition-delay: 0s')
          span M
        p(style='transition-delay: 0.02s')
          span A
        p(style='transition-delay: 0.04s')
          span N
        p(style='transition-delay: 0.06s')
          span A
        p(style='transition-delay: 0.08s')
          span G
        p(style='transition-delay: 0.10s')
          span E
        p(style='transition-delay: 0.12s')
          span M
        p(style='transition-delay: 0.14s')
          span E
        p(style='transition-delay: 0.16s')
          span N
        p(style='transition-delay: 0.18s')
          span T
      a.extension(v-if='isLogin', href='#', @click.stop.prevent='toWebrtc')
        p(style='transition-delay: 0s')
          span A
        p(style='transition-delay: 0.02s')
          span C
        p(style='transition-delay: 0.04s')
          span C
        p(style='transition-delay: 0.06s')
          span O
        p(style='transition-delay: 0.08s')
          span U
        p(style='transition-delay: 0.10s')
          span N
        p(style='transition-delay: 0.12s')
          span T
</template>

<script lang='ts'>
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mapState } from 'vuex';
import axios from 'axios';
import jQuery from 'jquery';
// tslint:disable-next-line:no-var-requires
const fs = require('fs');

@Component({
  computed: mapState({
    isLogin: 'isLogin',
    email: 'email',
    isAdmin: 'isAdmin',
  }),
})
export default class FooterNav extends Vue {
  @Prop() private msg!: string;

  private login(): void {
    this.$store.dispatch('login');
  }

  private logout(): void {
    this.$store.dispatch('logout');
  }

  private toHome(): void {
    this.$store.commit('setViewIndex', {
      index: 0,
    });
    this.$router.push({ name: 'home' });
  }

  private toManagement(): void {
    this.$router.push({ name: 'management-account', params: { user: 'admin' } });
  }

  private toWebrtc(): void {
    this.$router.push({ name: 'webrtc-account' });
  }
}
</script>

<style scoped lang='scss'>
$mes-color: #42b983;

#footer-nav {
  position: fixed;
  right: 0; bottom: 15px; left: 0;
  width: auto;
  height: auto;
}

a {
  font-size: 10px;
  color: $mes-color;
}

#user-account {
  font-size: 12px;
}

#footer-icon {
  height: 80px;
  cursor: pointer;
  animation: rotate_anime;
  animation-duration: 1s;
  opacity: .8;
  margin-right: 3px;
  display: none;
}

#footer-icon2 {
  width: 160px;
  display: none;
}

.extension {
  margin-left: 60px;
  margin-right: 60px;
  color: $mes-color;

  p {
    display: inline-block;
    transition: 0.5s;
    will-change: transform;
  }
}

.horizontal-array {
  display: flex;
  justify-content: center;

  img {
    height: 12px;
  }
  
  div {
    width: 15px;
    margin-left: 10px;
    margin-right: 10px;
    cursor: pointer;
  }
}

.icon.button {
  transition: 0.5s;
  opacity: .8;
}

.icon.button:hover {
  transform: rotateX( 360deg );
}

.extension:hover {
  p {
    -webkit-transform: rotateX(360deg);
    transform: rotateX(360deg);
  }
} 

@media screen and (max-width: 768px){
  a {
    font-size: 10px;
  }

  #user-account {
    font-size: 7px;
  }

  .extension {
    font-size: 8px;
  }
}

@keyframes rotate_anime {
    0% {
        transform:rotateX(0deg);
    }
    100% {
        transform:rotateX(360deg);
    }
}
</style>
