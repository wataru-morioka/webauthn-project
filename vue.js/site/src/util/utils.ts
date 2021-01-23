import axios from 'axios';
import firebase from 'firebase/app';
import { PhotoInfo, Article, ArticleInfo } from '../model/models';

class Result {
    public photoMultiArray: PhotoInfo[][];
    public projectTitleMap: Map<number, string>;

    constructor() {
      this.photoMultiArray = [];
      this.projectTitleMap = new Map<number, string>();
    }
}

// プロジェクト情報リスト取得
export const getPhotoList = () => {
    return new Promise<Result>(async (resolve, reject) => {
        const result = new Result();
        const photoMultiArray: PhotoInfo[][] = [];
        const projectTitleMap: Map<number, string> = new Map<number, string>();
        let photoArray: PhotoInfo[] = [];
        const currentUser = firebase.auth().currentUser!;
        const token = await currentUser.getIdToken(true);
        const header = {
            // Authorization: `Bearer ${token}`,
            // 'Access-Control-Allow-Origin': '*',
            // 'Access-Control-Allow-Headers': 'X-Requested-With, Origin, X-Csrftoken, Content-Type, Accept',
            // 'Access-Control-Allow-Methods': 'GET, PUT, POST, DELETE, OPTIONS',
            // 'Access-Control-Allow-Credentials': 'true',
        };

        // location.replace('http://localhost:8000/login');

        // await axios.get('https://www.google.com', {
        //     // headers: header,
        //     // withCredentials: true
        // })
        // .then((res) => {
        //     alert(res.data.message);
        //     resolve(result);
        //     return;
        // })
        // .catch((err) => {
        //     console.log('test');
        //     console.log(err);
        //     alert('test');
        //     reject();
        //     return;
        // });


        // // ヘッダに認証用トークンをセット
        // await axios.get('https://express.management/photographs', {
        //     headers: header,
        // })
        // .then((res) => {
        //     if (!res.data.result) {
        //         console.log('photoリスト取得に失敗しました');
        //         reject();
        //         return;
        //     }
        //     console.log('photoリスト取得');
        //     const resArray = res.data.photoArray;
        //     let photo: PhotoInfo;
        //     let index = 0;

        //     // プロジェクト一覧画面にて左右に1つずつ並べるため、2つで1組みの配列に格納
        //     resArray.forEach((el: any) => {
        //         photo = new PhotoInfo(
        //             el.id,
        //             el.sub_title,
        //             el.title,
        //             el.mimetype,
        //             el.fileName,
        //             el.size,
        //             el.data,
        //             el.created_datetime,
        //             el.modified_datetime,
        //         );

        //         const map = new Map<number, string>();
        //         projectTitleMap.set(photo.id, photo.title);
        //         photoArray.push(photo);

        //         if (index === 1) {
        //             photoMultiArray.push(photoArray);
        //             photoArray = [];
        //             index = 0;
        //             return;
        //         }

        //         index++;
        //     });

        //     if (resArray.length % 2 === 1) {
        //     photoMultiArray.push(photoArray);
        //     }

        //     result.photoMultiArray = photoMultiArray;
        //     result.projectTitleMap = projectTitleMap;
        //     resolve(result);
        // })
        // .catch((err) => {
        //     console.log('photoリスト取得に失敗しました');
        //     reject();
        // });
    });
};

// 記事リスト（コメントを外部結合した状態）取得（3件ずつ）
export const getArticleList = (currentArticleId: number, additional: boolean) => {
    return new Promise<ArticleInfo[]>(async (resolve, reject) => {
        const articleArray: ArticleInfo[] = [];
        const currentUser = firebase.auth().currentUser!;
        const token = await currentUser.getIdToken(true);
        const header = {
            Authorization: `Bearer ${token}`,
        };

        // ヘッダに認証用トークンをセット
        await axios.get('https://django.service/api/service/article', {
            headers: header,
            params: {
                current_article_id: currentArticleId,
                additional_flag: additional,
            },
        })
        .then((res) => {
            if (!res.data.result) {
                console.log('articleリスト取得に失敗しました');
                reject();
                return;
            }
            console.log('articleリスト取得');

            const resArray = res.data.articleList;
            let article: ArticleInfo;

            resArray.forEach((el: any) => {
                // アカウントのサムネイルデータはbase64形式で取得するので、バイト配列に変換
                let thumbnailBase64 = el.thumbnail;
                let buffer = new Uint8Array();

                if (thumbnailBase64 !== null) {
                    const bin = atob(thumbnailBase64.replace(/^.*,/, ''));
                    buffer = new Uint8Array(bin.length);
                    for (let i = 0; i < bin.length; i++) {
                    buffer[i] = bin.charCodeAt(i);
                    }
                }

                thumbnailBase64 = el.commentator_thumbnail;
                let commentBuffer = new Uint8Array();
                if (thumbnailBase64 !== null) {
                    const bin = atob(thumbnailBase64.replace(/^.*,/, ''));
                    commentBuffer = new Uint8Array(bin.length);
                    for (let i = 0; i < bin.length; i++) {
                    commentBuffer[i] = bin.charCodeAt(i);
                    }
                }

                article = new ArticleInfo(
                    el.id,
                    el.contributor_name,
                    el.body,
                    buffer,
                    el.created_datetime,
                    el.modified_datetime,
                    el.commentator_name,
                    commentBuffer,
                    el.comment_body,
                    el.comment_created_datetime,
                );
                articleArray.push(article);
            });

            resolve(articleArray);
        })
        .catch((err) => {
            console.log('articleリスト取得に失敗しました');
            reject();
        });
    });
};

// コメントが複数ある記事があるので、記事ごとにdistinct（key:記事ID、value:記事情報）
export const getDistinctArticleMap = (articleList: ArticleInfo[]): Map<number, Article> => {
    return new Map<number, Article>(
      articleList.map((article) => [
          article.id,
          new Article(
            article.id,
            article.contributorName,
            article.body,
            article.thumbnail,
            article.createdDatetime,
            article.modifiedDatetime,
          ),
        ],
      ),
    );
};

export const states = new Array<string>(
    '北海道',
    '青森県',
    '岩手県',
    '宮城県',
    '秋田県',
    '山形県',
    '福島県',
    '茨城県',
    '栃木県',
    '群馬県',
    '埼玉県',
    '千葉県',
    '東京都',
    '神奈川県',
    '新潟県',
    '富山県',
    '石川県',
    '福井県',
    '山梨県',
    '長野県',
    '岐阜県',
    '静岡県',
    '愛知県',
    '三重県',
    '滋賀県',
    '京都府',
    '大阪府',
    '兵庫県',
    '奈良県',
    '和歌山県',
    '鳥取県',
    '島根県',
    '岡山県',
    '広島県',
    '山口県',
    '徳島県',
    '香川県',
    '愛媛県',
    '高知県',
    '福岡県',
    '佐賀県',
    '長崎県',
    '熊本県',
    '大分県',
    '宮崎県',
    '鹿児島県',
    '沖縄県',
    'Foreign Country',
);
