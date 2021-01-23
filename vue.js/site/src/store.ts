import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import firebase from 'firebase/app';
import { PhotoInfo, Article, ArticleInfo } from './model/models';
import { getPhotoList, getDistinctArticleMap, getArticleList } from './util/utils';
// tslint:disable-next-line:no-var-requires
const fs = require('fs');
// tslint:disable-next-line:no-var-requires
const https = require('https');

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    userName: 'wataru',
    isLogin: false,
    isAnonymous: false,
    isVip: false,
    isAdmin: false,
    uid: '',
    idToken: '',
    email: '',
    displayName: '',
    state: '',
    thumbnail: new ArrayBuffer(0),
    currentViewIndex: 0,
    photoMultiArray: new Array<PhotoInfo[]>(),
    projectTitleMap: new Map<number, string>(),
    articleArray: new Array<ArticleInfo>(),
    distinctArticleMap: new Map<number, Article>(),
    currentArticleId: 0,
    isVideoDisplay: false,
    isVideoPlaying: false,
    authHeader: {},
  },

  mutations: {
    setUser(state, payload) {
      state.userName = payload.name;
    },

    setDisplayName(state, payload) {
      state.displayName = payload.displayName;
    },

    setUserInfo(state, payload) {
      state.isLogin = payload.isLogin;
      state.isAnonymous = payload.isAnonymous;
      state.thumbnail = payload.thumbnail;
      state.state = payload.state;
      state.isVip = payload.isVip;
      state.isAdmin = payload.isAdmin;
    },

    setViewIndex(state, payload) {
      state.currentViewIndex = payload.index;
    },

    setPhotoMutiArray(state, payload) {
      state.photoMultiArray = payload.photoMultiArray;
      state.projectTitleMap = payload.projectTitleMap;
    },

    setInitVideoFlag(state, payload) {
      state.isVideoDisplay = false;
      state.isVideoPlaying = false;
    },

    setIsDisplay(state, payload) {
      state.isVideoDisplay = payload.isVideoDisplay;
    },

    setIsPlaying(state, payload) {
      state.isVideoPlaying = payload.isVideoPlaying;
    },

    setHeader(state, payload) {
      state.authHeader = payload.authHeader;
      state.displayName = payload.displayName;
      state.idToken = payload.idToken;
      state.uid = payload.uid;
      state.email = payload.email;
    },

    setArticleArray(state, payload) {
      if (state.currentArticleId === 0 || !payload.additionalFlag) {
        state.distinctArticleMap = payload.distinctArticleMap;
      } else {
        // スクロールして追加記事を取得した場合、Mapに追加
        Array.from(payload.distinctArticleMap as Map<number, ArticleInfo>).map(([key, value]) => {
          state.distinctArticleMap.set(key, value);
        });
      }

      if (!payload.additionalFlag) {
        state.articleArray = payload.articleArray;
      } else {
        // スクロールして追加記事を取得した場合、配列に追加
        state.articleArray = state.articleArray.concat(payload.articleArray);
      }

      state.currentArticleId = payload.currentArticleId;
    },
  },

  actions: {
    // プロジェクト情報リスト取得
    async getPhotos({ commit, state, rootState }) {
      await getPhotoList().then((result) => {
        this.commit('setPhotoMutiArray', {
          photoMultiArray: result.photoMultiArray,
          projectTitleMap: result.projectTitleMap,
        });
      }).catch((err) => {
        // alert('err');
      });
    },

    // ブログ記事（コメントを左外部結合した状態）リスト（3件）取得
    async getArticles({ commit, state, rootState }, { additional }) {
      await getArticleList(state.currentArticleId, additional).then((articleList) => {
        if (articleList.length === 0) {
          alert('これ以上記事はありません');
          return;
        }
        const distinctArticleMaps = getDistinctArticleMap(articleList);
        const maxIdArticle = articleList.reduce((a, b) => {
          return a.id > b.id ? b : a;
        });
        this.commit('setArticleArray', {
          additionalFlag: additional,
          currentArticleId: maxIdArticle.id,
          articleArray: articleList,
          distinctArticleMap: distinctArticleMaps,
        });
      }).catch((err) => {
        // alert('err');
      });
    },

    // home画面「Login」ボタンを押した時
    async login({ commit, state, rootState }) {
      const provider = new firebase.auth.GoogleAuthProvider();
      console.log(provider);
      await firebase.auth().signInWithPopup(provider)
      .then(async (result) => {
        console.log('google認証');

        // DBにユーザ登録
        await this.dispatch('getHeader');
        const body = {
          name: state.displayName,
        };

        // ヘッダに認証用トークンをセット
        await axios.post('https://django.service:443/api/service/account', body, {
          headers: state.authHeader,
        })
        .then((res) => {
          if (!res.data.result) {
            console.log('サーバのログイン処理に失敗しました');
          }
          console.log('ログインしました');
        })
        .catch((err) => {
          console.log('サーバのログイン処理に失敗しました');
        });
      }).catch((error) => {
        console.log(error);
        console.log('ログインに失敗しました');
      });
    },

    // home画面「Logout」ボタンを押した時
    async logout({ commit, state, rootState }) {
      await firebase.auth().signOut()
      .then((result) => {
        console.log('ログアウト');
      })
      .catch((err) => {
        console.log('ログアウトに失敗しました');
      });
    },

    // このサイトに訪れる度に実行
    checkLoginStatus({ commit, state, rootState }) {
      firebase.auth().onAuthStateChanged( async (user) => {
        if (user) {
          this.dispatch('getPhotos');
          this.dispatch('getArticles', {
            additional: false,
          });
          await this.dispatch('getHeader');

          // 一度でもgoogle認証をしていた（ログインしていた）場合
          if (user.email != null && user.email.length > 0) {
            console.log('google認証済み');
            const body = {};

            // ヘッダに認証用トークンをセット
            await axios.put('https://django.service:443/api/service/account', body, {
              headers: state.authHeader,
            })
            .then((res) => {
              if (!res.data.result) {
                console.log('サーバのログイン処理に失敗しました');
                return;
              }

              console.log('ログイン');
              const thumbnailBase64 = res.data.thumbnail;
              let buffer = null;

              // ユーザがログイン後アカウント登録を完了していた場合
              if (thumbnailBase64 !== null) {
                // サムネイルはbase64形式で取得するので、バイト配列に変換
                const bin = atob(thumbnailBase64.replace(/^.*,/, ''));
                buffer = new Uint8Array(bin.length);
                for (let i = 0; i < bin.length; i++) {
                    buffer[i] = bin.charCodeAt(i);
                }
              }

              // ユーザ情報セット
              this.dispatch('setUserInfo', {
                isLoginAuth: true,
                isAnonymousAuth: false,
                thumbnailData: buffer,
                region: res.data.state,
                isVipAccount: res.data.isVip,
                isAdminAccount: res.data.isAdmin,
              });
            })
            .catch((err) => {
              console.log('サーバのログイン処理に失敗しました');
            });

            return;
          }
        // サイトの初訪問者
        } else {
          console.log('匿名認証');
          // firebase匿名認証
          firebase.auth().signInAnonymously().catch((error) => {
            const errorCode = error.code;
            const errorMessage = error.message;
            console.log(error.message);
          });
        }

        // このサイトに一度は訪問したことがあるが、まだログインしていない場合
        console.log('匿名認証済み');
        this.dispatch('setUserInfo', {
          isLoginAuth: false,
          isAnonymousAuth: true,
          thumbnailData: null,
          region: '',
          isVipAccount: false,
          isAdminAccount: false,
        });
      });
    },

    // ユーザ情報をセット
    async setUserInfo({ commit, state, rootState },
                      { isLoginAuth, isAnonymousAuth, thumbnailData, region, isVipAccount, isAdminAccount } ) {
      const currentUser = firebase.auth().currentUser;
      if (currentUser == null) {
        return;
      }
      const token = await currentUser.getIdToken(true).catch((err) => {
        console.log(err);
      });
      this.commit('setUserInfo', {
        isLogin: isLoginAuth,
        isAnonymous: isAnonymousAuth,
        thumbnail: thumbnailData,
        state: region,
        isVip: isVipAccount,
        isAdmin: isAdminAccount,
      });
    },

    // サーバに送信する用の認証ヘッダをセット
    async getHeader({ commit, state, rootState }) {
      const currentUser = firebase.auth().currentUser;
      if (currentUser == null) {
        return;
      }
      const token = await currentUser.getIdToken(true).catch((err) => {
        console.log(err);
      });
      const header = {
        Authorization: `Bearer ${token}`,
      };

      this.commit('setHeader', {
        authHeader: header,
        displayName: currentUser.displayName,
        idToken: token,
        uid: currentUser.uid,
        email: currentUser.email,
      });
    },
  },

  getters: {
    getName: (state, getters) => {
      return state.userName;
    },
    isLogin: (state, getters) => {
      return state.isLogin;
    },
    getPhotos: (state, getters) => {
      return state.photoMultiArray;
    },
    getProjectTitleMap: (state, getter) => {
      return state.projectTitleMap;
    },
    isVideoDisplay: (state, getters) => {
      return state.isVideoDisplay;
    },
    isVideoPlaying: (state, getters) => {
      return state.isVideoPlaying;
    },
  },
});
