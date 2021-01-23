<template lang='pug'>
  div#sub-menu(style='transition-delay: 3s')
    div.menu-item(v-for='[id, title] in this.$store.state.projectTitleMap', :key='id')
      a(href='#', @click.stop.prevent='play(id)')
        p(v-for='(char, charIndex) in Array.from(title)', :key='charIndex', :style='transitionDelay(0, 0.02, charIndex)')
          span {{ char }}
</template>

<script lang='ts'>
import Vue from 'vue';
import { Component, Mixin, Mixins } from 'vue-mixin-decorator';
import jQuery from 'jQuery';
import axios from 'axios';
import firebase from 'firebase/app';
import VideoMixin from '@/components/VideoMixin.vue';

@Component
export default class SubMenu extends Mixins<VideoMixin>(VideoMixin) {
  private fadein() {
    $('#sub-menu').css({
      opacity: 1,
      transform: 'translate(0px, 0px)',
    });
  }

  private async mounted() {
    this.fadein();
  }

  private transitionDelay(diff: number, rate: number, index: number): string {
    const delay = diff + rate * index;
    return 'transition-delay: ' + delay + 's';
  }
}
</script>

<style scoped lang='scss'>
#sub-menu {
  display: inline-block;
  height: 270px;
  width: 100px;
  opacity: 0;
  transform: translate(0px, 50px);
  transition: 2s;

  #update-menu-bt {
    display: none;
  }

  .menu-item {
    a p {
      transition: 0.6s;
      will-change: transform;
    }
  }

  .menu-item:hover {
    a p {
      -webkit-transform: rotateX(360deg);
      transform: rotateX(360deg);
    }
  }

  a {
    cursor: pointer;
  }

  p {
    cursor: pointer;
    display: inline-block;
    font-size: 11px;
    color: #ffffff8a;
  }
}
</style>
