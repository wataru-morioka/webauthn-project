<template lang='pug'>
  div#center-nav
    div#center
      div.icon-wrap(@click='toPrev')
        img.slide-icon.left(src='../assets/prev.png')
      div.icon-wrap(@click='toNext')
        img.slide-icon(src='../assets/next.png')
</template>

<script lang='ts'>
import { Component, Prop, Vue } from 'vue-property-decorator';
import axios from 'axios';
// tslint:disable-next-line:no-var-requires
const fs = require('fs');

@Component
export default class CenterNav extends Vue {
  private viewMap: { [key: number]: string; } = {
    0: 'home',
    1: 'sauza',
    2: 'jagermeister',
  };

  private toNext(): void {
    const targetIndex = this.getTargetIndex(1);
    this.$router.push({ name: this.viewMap[targetIndex] });
  }

  private toPrev(): void {
    const targetIndex = this.getTargetIndex(2);
    this.$router.push({ name: this.viewMap[targetIndex] });
  }

  private getTargetIndex(additional: number): number {
    const currentIndex = (this.$store.state.currentViewIndex + additional) % 3;
    this.$store.commit('setViewIndex', {
      index: currentIndex,
    });
    return currentIndex;
  }
}
</script>

<style scoped lang='scss'>
$menu-color: #ffffff;

#center-nav {
  position: absolute;
  width: 100%;
  top: 0px;
  right: 0px;
  bottom: 0px;
  left: 0px;
  margin: auto;
  height: 50px;
}

#center {
  display: flex;
  justify-content:space-between;
  -webkit-box-pack:justify;
}

.icon-wrap {
  cursor: pointer;
  margin-left: 15px;
  margin-right: 15px;
  transition: 0.7s ;
  will-change: transform;
  .left {
    opacity: 0;
  }
}

.icon-wrap:hover {
  transform: rotateX( 360deg );
}

.slide-icon {
  width: 15px;
}
</style>