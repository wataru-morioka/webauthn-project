<template lang='pug'>
  div
    div#bg-img 
      img.pc(src='../assets/bg-load.jpg')
      img.sp(src='../assets/bg-sp.jpg')
    div#bg-circle
      img#circle-img(src='../assets/bg-circle.png')
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class Back extends Vue {}
</script>

<style scoped lang='scss'>
#bg-img {
  img {
    position: fixed;
    width: 100%;
    top: 0;
    left: 0;
    overflow: hidden;
    min-width: 100%;
    min-height: 100%;
    z-index: -5;
  }
} 

#circle-img {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  width: 430px;
  z-index: -4;
}

@media screen and (max-width: 768px){
  .pc {
    display:none;
  }
  .sp {
    display:block;
  }
  #circle-img {
    width: 300px;
  }
}

@media screen and (min-width: 768px){
  .pc {
    display:block;
  }
  .sp {
    display:none;
  }
}
</style>
