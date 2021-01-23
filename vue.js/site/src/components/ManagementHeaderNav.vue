<template lang='pug'>
  div#header-nav
    div#management-sidebar(class="ui left sidebar inverted vertical menu")
      a.item(@click='toAccount') ACCOUNT
      a.item(@click='toUpload') UPLOAD
      a.item(@click='toContact') CONTACT
    div#header
      div#management-dropdown(class="ui pointing link icon dropdown")
        img.icon#header-logo(src='../assets/jager-logo.png', @click='openSidebar')
      VipMenuNav
</template>

<script lang='ts'>
import { Component, Prop, Vue } from 'vue-property-decorator';
import VipMenuNav from '@/components/VipMenuNav.vue';
import axios from 'axios';
import jQuery from 'jquery';
// tslint:disable-next-line:no-var-requires
const fs = require('fs');

@Component({
  components: {
    VipMenuNav,
  },
})
export default class ManagementHeaderNav extends Vue {
  private isDroped: boolean = false;
  private isOpen: boolean = false;

  private mounted() {
    ($('#management-sidebar') as any).sidebar({
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
    ($('#management-sidebar') as any).sidebar('show');
  }

  private toAccount(): void {
    this.$router.push({ name: 'management-account' });
    ($('#management-sidebar') as any).sidebar('hide');
  }

  private toUpload(): void {
    this.$router.push({ name: 'management-upload' });
    ($('#management-sidebar') as any).sidebar('hide');
  }

  private toContact(): void {
    this.$router.push({ name: 'management-contact' });
    ($('#management-sidebar') as any).sidebar('hide');
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

#header-logo {
  margin-top: 7px;
  margin-left: 10px;
  height: 50px;
  cursor: pointer;
  opacity: .8;
}

#management-sidebar {
  .item {
    color: #ffffff77;
  }
}

.menu {
  background: (0, 0, 0, 0) !important;
  div {
    color: $menu-color !important;
  }

  .item {
    font-size: 12px !important;
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
