<template lang='pug'>
  div#header-nav
    div#webrtc-sidebar(class="ui left sidebar inverted vertical menu")
      a.item(v-if='isAdmin', @click='toArticle') ARTICLE
      a.item(@click='toAccount') ACCOUNT
    div#header
      div#webrtc-dropdown(class="ui pointing link icon dropdown")
        img.icon#header-logo(src='../assets/jager-logo.png', @click='openSidebar')
      VipMenuNav
</template>

<script lang='ts'>
import { Component, Prop, Vue } from 'vue-property-decorator';
import VipMenuNav from '@/components/VipMenuNav.vue';
import { mapState } from 'vuex';
import axios from 'axios';
import jQuery from 'jquery';
// tslint:disable-next-line:no-var-requires
const fs = require('fs');

@Component({
  components: {
    VipMenuNav,
  },
  computed: mapState({
    isAdmin: 'isAdmin',
  }),
})
export default class WebRtcHeaderNav extends Vue {
  private isDroped: boolean = false;
  private isOpen: boolean = false;

  private mounted() {
    ($('#webrtc-sidebar') as any).sidebar({
      onShow: () => {
        this.isOpen = true;
      },
      onHide: () => {
        this.isOpen = false;
      },
    })
    .sidebar('setting', 'transition', 'push')
    .sidebar('toggle');
  }

  private openSidebar(): void {
    ($('#webrtc-sidebar') as any).sidebar('show');
  }

  private toArticle(): void {
    this.$router.push({ name: 'webrtc-article' });
    ($('#webrtc-sidebar') as any).sidebar('hide');
  }

  private toAccount(): void {
    this.$router.push({ name: 'webrtc-account' });
    ($('#webrtc-sidebar') as any).sidebar('hide');
  }
}
</script>

<style scoped lang='scss'>
$menu-color: #ffffff77;

#header {
  display: flex;
  justify-content:space-between;
  -webkit-box-pack:justify;
}

#webrtc-sidebar {
  .item {
    color: #ffffff77;
  }
}

#header-logo {
  margin-top: 17px;
  margin-left: 50px;
  height: 50px;
  cursor: pointer;
  opacity: .8;
}

.item {
  font-size: 12px !important;
}

.menu {
  background: (0, 0, 0, 0) !important;
  div {
    color: $menu-color !important;
  }
}

@media screen and (max-width: 768px){
  #header-logo {
    margin-top: 12px;
    margin-left: 20px;
    height: 40px;
  }

  .item {
    font-size: 10px !important;
  }
}

@media screen and (min-width: 768px){
}
</style>
