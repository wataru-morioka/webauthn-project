<template lang='pug'>
  div#start
    button#to-home-bt(@click='toHome')
    div(class="ui indicating progress")
      div(class="bar")
      div(class="label")
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import jQuery from 'jquery';

window.onload = () => {
  const $progress = $('.ui.progress') as any;

  // // 初期ロードの時間かせぎにプログレスバーを表示
  // // TODO 何かしたのモーションを追加
  // setInterval(() => {
  //   $progress.progress('increment');
  //   if ($progress.progress('is complete')) {
  //     setTimeout(() => {
  //       const bt = document.getElementById('to-home-bt');
  //       if (bt == null) {
  //         return;
  //       }
  //       bt.click();
  //     }, 400);
  //   }
  // // }, 50);
  // }, 5);
};

@Component
export default class Start extends Vue {
  private mounted() {
    // クッキーが存在しない場合、もしくは初期イニシャルapiアクセス時、リフレッシュトークンが無効な場合
    location.replace('http://localhost:8000/login');


    // ($('.ui.progress') as any).progress({
    //   duration : 400,
    //   total    : 100,
    //   text     : {
    //     active: '{value} %',
    //   },
    // });
  }

  private toHome(): void {
    
    this.$router.push({ name: 'home', params: { user: 'admin' } });
  }
}
</script>

<style scoped lang='scss'>
#start {
  width: 100%;
  height: 100vh;
  position: relative;
  overflow: hidden;
  background: rgba(0, 0, 0, .9);
  animation-name: body-fadein;
  animation-duration: 2s;
}

.ui.progress {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  width: 200px;
  height:1px;
}

.ui.progress .bar {
  height: 2px;
}

.ui.indicating .label {
  color: #ffffff77 !important;
  font-size: 8px;
  font-family: 'Avenir', Helvetica, Arial, sans-serif !important;
}

#to-home-bt {
  display: none;
}
</style>
