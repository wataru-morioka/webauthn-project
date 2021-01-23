<template lang='pug'>
  div#header-nav
    div#header
      div
        img#header-logo(src='../assets/jagermeister.png', @click='toHome')
      div#header-menu
        div
          div.rotate-menu
            a(href='#', @click.stop.prevent='grapqhqlRequest')
              p(style='transition-delay: 0s')
                span G
              p(style='transition-delay: 0.02s')
                span R
              p(style='transition-delay: 0.04s')
                span A
              p(style='transition-delay: 0.06s')
                span P
              p(style='transition-delay: 0.08s')
                span H
              p(style='transition-delay: 0.10s')
                span Q
              p(style='transition-delay: 0.12s')
              span L
          div.rotate-menu
            a(href='#', @click.stop.prevent='toHome')
              p(style='transition-delay: 0s')
                span H
              p(style='transition-delay: 0.02s')
                span O
              p(style='transition-delay: 0.04s')
                span M
              p(style='transition-delay: 0.06s')
                span E
          div.rotate-menu
            a(href='#', @click.stop.prevent='toAbout')
              p(style='transition-delay: 0s')
                span A
              p(style='transition-delay: 0.02s')
                span B
              p(style='transition-delay: 0.04s')
                span O
              p(style='transition-delay: 0.06s')
                span U
              p(style='transition-delay: 0.08s')
                span T
          div.rotate-menu
            a(href='#', @click.stop.prevent='toProject')
              p(style='transition-delay: 0s')
                span P
              p(style='transition-delay: 0.02s')
                span R
              p(style='transition-delay: 0.04s')
                span O
              p(style='transition-delay: 0.06s')
                span J
              p(style='transition-delay: 0.08s')
                span E
              p(style='transition-delay: 0.10s')
                span C
              p(style='transition-delay: 0.12s')
                span T
          div.rotate-menu
            a(href='#', @click.stop.prevent='toArticle')
              p(style='transition-delay: 0s')
                span A
              p(style='transition-delay: 0.02s')
                span R
              p(style='transition-delay: 0.04s')
                span T
              p(style='transition-delay: 0.06s')
                span I
              p(style='transition-delay: 0.08s')
                span C
              p(style='transition-delay: 0.10s')
                span L
              p(style='transition-delay: 0.12s')
                span E
          div.rotate-menu
            a(href='#', @click.stop.prevent='toMember')
              p(style='transition-delay: 0s')
                span M
              p(style='transition-delay: 0.02s')
                span E
              p(style='transition-delay: 0.04s')
                span M
              p(style='transition-delay: 0.06s')
                span B
              p(style='transition-delay: 0.08s')
                span E
              p(style='transition-delay: 0.10s')
                span R
          div.rotate-menu
            a(href='#', @click.stop.prevent='toContact')
              p(style='transition-delay: 0s')
                span C
              p(style='transition-delay: 0.02s')
                span O
              p(style='transition-delay: 0.04s')
                span N
              p(style='transition-delay: 0.06s')
                span T
              p(style='transition-delay: 0.08s') 
                span A
              p(style='transition-delay: 0.10s')
                span C
              p(style='transition-delay: 0.12s')
                span T
        div#menu-dropdown(class="ui pointing link icon dropdown")
          i(class="large bars icon")
          div.menu
            div.header MENU
            div.divider
            div.item(@click='toHome') HOME
            div.item(@click='toAbout') ABOUT
            div.item(@click='toProject') PROJECT
            div.item(@click='toArticle') ARTICLE
            div.item(@click='toMember') MEMBER
            div.item(@click='toContact') CONTACT
</template>

<script lang='ts'>
import { Component, Prop, Vue } from 'vue-property-decorator';
import { ALL_POST_QUERY } from './../graphql/operations';
import { useQuery } from '@vue/apollo-composable';
import { OrderBy } from "./../graphql/graphql-types";
import { GetAllPosts } from "./../graphql/interface/GetAllPosts";
import { ref } from '@vue/composition-api'
// tslint:disable-next-line:no-var-requires
const fs = require('fs');

@Component
export default class HeaderNav extends Vue {
  private isDroped: boolean = false;

  private mounted() {
    // メニューアイコンを押下した際、ドロップダウンメニューが下に伸びるため、コンテンツ部分のzインデックスを操作
    ($('#menu-dropdown') as any).dropdown({
      on: 'hover',
      onShow: () => {
        $('.content, .contents').css('z-index', '-3');
      },
      onHide: () => {
        $('.content, .contents').css('z-index', '0');
      },
    });
  }

  private grapqhqlRequest(): void {
    const { result, loading, error } = useQuery<GetAllPosts>(ALL_POST_QUERY, 
      ref ({ 
        orderBy: OrderBy.createdAt_ASC,
        first: 1,
        skip: 3
      })
    );

    console.log(result);
  }

  private toHome(): void {
    this.$store.commit('setViewIndex', {
      index: 0,
    });
    this.$router.push({ name: 'home' });
  }

  private toAbout(): void {
    this.$router.push({ name: 'about' });
  }

  private toProject(): void {
    this.$router.push({ name: 'project' });
  }

  private toArticle(): void {
    this.$router.push({ name: 'article' });
  }

  private toMember(): void {
    this.$router.push({ name: 'member' });
  }

  private toContact(): void {
    this.$router.push({ name: 'contact' });
  }
}
</script>

<style scoped lang='scss'>
$menu-color: #ffffffbf;

#header {
  display: flex;
  justify-content:space-between;
  -webkit-box-pack:justify;
}

#header-logo {
  margin-top: 10px;
  margin-left: 50px;
  height: 40px;
  cursor: pointer;
  opacity: .8;
}

#header-menu {
  margin-top: 20px;
  display: flex;
  justify-content:space-between;
  -webkit-box-pack:justify;
  position: relative;
}

a {
  font-size: 9px;
  color: $menu-color;
  p {
    color: $menu-color;
  }
}

#temp {
  width: 100px;
}

.menu {
  background: (0, 0, 0, 0) !important;
  div {
    color: $menu-color !important;
  }
}

@media screen and (max-width: 768px){
  .rotate-menu {
    display:none;
  }

  #header-logo {
    margin-top: 20px;
    margin-left: 20px;
    height: 20px;
  }

  #header-menu {
    margin-top: 20px;
  }

  #menu-dropdown {
    .item {
      font-size: 10px;
    }
  }

  .bars {
    margin-right: 20px !important;
  }
}

@media screen and (min-width: 768px){
  .rotate-menu {
    margin-right: 80px;
    display: inline-block;
  }
}

a {
  p {
    display: inline-block;
    transition: 0.5s;
    will-change: transform;
  }
} 

a:hover {
  p {
    -webkit-transform: rotateX(360deg);
    transform: rotateX(360deg);
  }
} 

.bars {
  margin-right: 50px;
  cursor: pointer;
}

@keyframes rotate_anime {
    0% {
        transform: rotateX(0deg);
    }
    100% {
        transform: rotateX(360deg);
    }
}
</style>
