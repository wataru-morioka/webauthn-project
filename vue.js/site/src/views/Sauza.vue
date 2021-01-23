<template lang='pug'>
  div#sauza
    video#sauza-video(muted playsinline)
      //- source(src='../assets/sauza.mp4' type='video/mp4')
      //- source(src='' type='video/mp4')
    MainNav
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import MainNav from '@/components/MainNav.vue';
import Hls from 'hls.js';

@Component({
  components: {
    MainNav,
  },
})
export default class Sauza extends Vue {
  private mounted() {

    // hls.js（HTTP Live Streaming）を利用し、動画をダウンロードしながら再生
    const video = document.getElementById('sauza-video') as HTMLVideoElement;
    if (Hls.isSupported()) {
      const hls = new Hls();
      hls.loadSource('https://dggt9dial88dg.cloudfront.net/hls/sauza/index.m3u8');
      hls.attachMedia(video);
      hls.on(Hls.Events.MANIFEST_PARSED, () => {
        video.play();
      });
    } else if (video.canPlayType('application/vnd.apple.mpegurl')) {
      video.src = 'https://dggt9dial88dg.cloudfront.net/hls/sauza/index.m3u8';
      video.addEventListener('loadedmetadata', () => {
          video.play();
      });
    }
  }
}
</script>

<style scoped lang='scss'>
#sauza {
  width: 100%;
  height: 100vh;
  position: relative;
  overflow: hidden;
  background: rgba(0, 0, 0, .8);
  animation-duration: 1s;
}
 
#sauza-video {
  position: absolute;
  top: 0;
  left: 0;
  width: auto;
  height: auto;
  min-width: 100%;
  min-height: 100%;
  z-index: -1;
}
</style>
