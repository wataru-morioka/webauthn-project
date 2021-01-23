<template lang='pug'>
  div.content(class="ui two column divided grid", @click='stop')
    Modal
    div.row(v-for='(photoArray, index) in this.$store.state.photoMultiArray', :key='index', :style='firstChildRow(index)')
      div.column.left(:style='transitionDelay(0, 0.2, index * 2)')
        div.content-box.box-left
          div.image(@click='play(photoArray[0].id)')
            i.huge.play.icon
            img(:src='setPhoto(index, 0)')
          div.subject
            p {{ photoArray[0].subTitle }}
            h4 
              p(v-for='(char, charIndex) in Array.from(photoArray[0].title)', :key='charIndex', :style='transitionDelay(0, 0.02, charIndex)')
                span {{ char }}
      div.column(v-if='existsPhoto(photoArray[1])', :style='transitionDelay(0, 0.2, index * 2 + 1)')
        div.content-box.box-right
          div.image(@click='play(photoArray[1].id)')
            i.huge.play.icon
            img(:src='setPhoto(index, 1)')
          div.subject
            p {{ photoArray[1].subTitle }}
            h4 
              p(v-for='(char, charIndex) in Array.from(photoArray[1].title)', :key='charIndex', :style='transitionDelay(0, 0.02, charIndex)')
                span {{ char }}
</template>

<script lang='ts'>
import Vue from 'vue';
import { Component, Mixin, Mixins } from 'vue-mixin-decorator';
import Modal from '@/components/Modal.vue';
import VideoMixin from '@/components/VideoMixin.vue';
import axios from 'axios';
import jQuery from 'jQuery';
import { PhotoInfo } from '../model/models';

@Component({
  components: {
    Modal,
  },
})
export default class ProjectContent extends Mixins<VideoMixin>(VideoMixin) {
  private firstChildRow(index: number): string {
    if (index === 0) {
      return 'margin-top: 70px;';
    }
    return 'margin-top: auto;';
  }

  private existsPhoto(photo: any): boolean {
    if (photo === undefined) {
      return false;
    }
    return true;
  }

  private setPhoto(indexX: number, indexY: number): string {
    const buffer = Buffer.from(this.$store.state.photoMultiArray[indexX][indexY].data);
    const blob = new Blob([buffer], {type: this.$store.state.photoMultiArray[indexX][indexY].mimetype});
    const blobURL = window.URL.createObjectURL(blob);
    return blobURL;
  }

  private mounted() {
    setTimeout(() => {
      this.fadein();
    }, 2500);
    $('.content').scroll(() => {
      this.fadein();
      (this as any).stop();
    });
  }

  private transitionDelay(diff: number, rate: number, index: number): string {
    const delay = diff + rate * index;
    return 'transition-delay: ' + delay + 's';
  }

  private fadein(): void {
    const offset = 200;
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
}
</script>

<style scoped lang='scss'>
.content {
  position: fixed;
  top: 60px; right: 0; bottom: 0; left: 0;
  overflow-y: scroll;
  -webkit-overflow-scrolling: touch;
  height: auto;
  margin: auto;
  transition: 1s;
  will-change: transform;
}

::-webkit-scrollbar {
  display: none;
  -webkit-appearance: none;
}

.image {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  i {
    position: fixed;
  }
}

img {
  width: 100%;
}
  
.row {
  margin-top: -30px;
}

.box-left {
  margin-left: auto;
  margin-right: 150px;
}

.box-right {
  margin-top: 70px;
  margin-left: 150px;
  margin-bottom: 30px;
}
  
.column {
  min-width: 50%;
  width: 100%;
  opacity: 0;
  transform: translate(0px, 30px) translate3d(0, 0, 0);
  transition: 1s;
  will-change: transform;
  -ms-filter: blur(30px);
  filter: blur(30px);
  
  .content-box {
    width: 340px;
    cursor: pointer;

    .subject {
      text-align: left;
      margin-top: 10px;
      p {
        font-size: 12px;
        margin-bottom: 0px;
      }
      h4 {
        margin-top: 0px;
        color: #ffffff;
        p {
          font-size: 16px;
          display: inline-block;
          transition: 0.6s
        }
      }
    }
  }
}

.content-box:hover {
  h4 p {
    -webkit-transform: rotateX(360deg);
    transform: rotateX(360deg);
  }
}

@media screen and (max-width: 768px){
  .content {
    .row {
      margin: auto !important;
      padding-top: 0px;
    }
  }

  .box-left {
    margin: auto;
  }

  .box-right {
    margin: auto;
  }

  .column {
    min-width: 100%;
    margin-bottom: 30px !important;
    .content-box {
      width: 80%;

      .subject {
        margin-top: 0px;

        p {
          font-size: 6px;
        }

        h4 p {
          padding-top: 0px;
          font-size: 13px;
        }
      }
    }
  }

  .column.left {
    margin-bottom: 60px !important;
  }
}
</style>
