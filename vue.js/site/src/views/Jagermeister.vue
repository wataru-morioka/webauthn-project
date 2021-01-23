<template lang='pug'>
  div#jagermeister
    video#jager-video(muted playsinline)
      //- source(src='../assets/jager.mp4' type='video/mp4')
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
export default class Jagermeister extends Vue {
  private mounted() {
    const video = document.getElementById('jager-video') as HTMLVideoElement;

    // hls.js（HTTP Live Streaming）を利用し、動画をダウンロードしながら再生
    if (Hls.isSupported()) {
      const hls = new Hls();
      hls.loadSource('https://dggt9dial88dg.cloudfront.net/hls/jager/index.m3u8');
      hls.attachMedia(video);
      hls.on(Hls.Events.MANIFEST_PARSED, () => {
        video.play();
      });
    } else if (video.canPlayType('application/vnd.apple.mpegurl')) {
      video.src = 'https://dggt9dial88dg.cloudfront.net/hls/jager/index.m3u8';
      video.addEventListener('loadedmetadata', () => {
          video.play();
      });
    }
  }
}
</script>

<style scoped lang='scss'>
#jagermeister {
  width: 100%;
  height: 100vh;
  position: relative;
  overflow: hidden;
  background: rgba(0, 0, 0, .8);
  animation-duration: 1s;
}
 
#jager-video {
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