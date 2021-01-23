<template lang='pug'>
  div.content
    div(class="ui two column divided grid")
      div.row#mimber-1
        div.column.left(style='transition-delay: 2s;')
          img(src='../assets/member.jpg')
        div.column.right(style='transition-delay: 2.5s;')
          h5 Director
          table.profile
            tr
              td.title
                p name
              td.value
                p Watau Morioka
            tr
              td.title
                p skill
              td.value
                p direction photo 
            tr
              td.title
                p hobby
              td.value
                p 釣り  音楽鑑賞  
            tr
              td.title
                p phone
              td.value
                p 090-0000-0000
            tr
              td.title
                p email
              td.value
                p example@gmail.com
          br
          img.icon(src='../assets/icn_fb.png')
          img.icon(src='../assets/icn_insta.png')
          img.icon(src='../assets/icn_tw.png')
      div.row#mimber-2
        div.column.left(style='transition-delay: 2.2s;')
          img(src='../assets/member.jpg')
        div.column.right(style='transition-delay: 2.7s;')
          h5 Video Photographer
          table.profile
            tr
              td.title
                p name
              td.value
                p Watau Morioka
            tr
              td.title
                p skill
              td.value
                p video photo edit
            tr
              td.title
                p hobby
              td.value
                p ライブ  飲み歩き  映画鑑賞  
            tr
              td.title
                p phone
              td.value
                p 090-0000-0000
            tr
              td.title
                p email
              td.value
                p example@gmail.com
          br
          img.icon(src='../assets/icn_fb.png')
          img.icon(src='../assets/icn_insta.png')
          img.icon(src='../assets/icn_tw.png')
      div.row#mimber-3
        div.column.left(style='transition-delay: 2.4s;')
          img(src='../assets/member.jpg')
        div.column.right(style='transition-delay: 2.9s;')
          h5 Editor
          table.profile
            tr
              td.title
                p name
              td.value
                p Watau Morioka
            tr
              td.title
                p skill
              td.value
                p edit video 
            tr
              td.title
                p hobby
              td.value
                p 旅行  フットサル  読書  
            tr
              td.title
                p phone
              td.value
                p 090-0000-0000
            tr
              td.title
                p email
              td.value
                p example@gmail.com
          br
          img.icon(src='../assets/icn_fb.png')
          img.icon(src='../assets/icn_insta.png')
          img.icon(src='../assets/icn_tw.png')
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import jQuery from 'jQuery';

@Component
export default class MemberContent extends Vue {
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
}
</script>

<style scoped lang='scss'>
.content {
  position: fixed;
  overflow-y: scroll;
  overflow-x: hidden;
  -webkit-overflow-scrolling: touch;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin: auto;
  height: auto;
  width: 100%;
  top: 60px; right: 0; bottom: 0; left: 0;
  transition: 1s;
}

::-webkit-scrollbar {
  display: none;
  -webkit-appearance: none;
}

.row {
  margin-top: 30px;
  margin-bottom: 30px;
}

#mimber-1 {
  margin-right: 300px;
}

#mimber-3 {
  margin-left: 300px;
}

.icon {
  cursor: pointer;
  height: 12px;
  margin-left: 10px;
  margin-right: 10px;
}

.column {
  min-width: 50%;
  opacity: 0;
  transform: translate(5px, 40px) translate3d(0, 0, 0);
  transition: 1s;
  -ms-filter: blur(50px);
  filter: blur(50px);

  .left {
    text-align: right;
  }

  .right {
    text-align: left;
  }
}

.profile tr {
  text-align: left;
  height: 20px;

  .title {
    width: 50px;
    font-size: 12px;
  }

  .value {
    padding-left: 50px;
    color: #ffffff;
    font-size: 12px;
  }
}

@media screen and (max-width: 768px){
  .content {
    justify-content: flex-start;
  }

  #mimber-1 {
    margin-right: auto;
  }

  #mimber-3 {
    margin-left: auto;
  }

  .column {
    min-width: 100%;
    .left {
      text-align: center;
      img {
        width: 70%;
      }
    }
    .right {
      text-align: center;
      h5 {
        font-size: 11px;
      }
    }
  }

  .row {
    table {
      margin: auto;
    }
  }

  .profile tr {
    .title {
      font-size: 9px;
    }

    .value {
      font-size: 9px;
    }
  }
}
</style>
