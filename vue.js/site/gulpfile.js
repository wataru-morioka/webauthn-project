// gulpプラグインを読み込みます
const gulp = require("gulp"),
    imgmin = require('gulp-imagemin'),
    imgminPNG = require('imagemin-pngquant'),
    imgminJPG = require('imagemin-jpeg-recompress'),
    imgminGIF = require('imagemin-gifsicle'),
    imgminSVG = require('imagemin-svgo');

const minifyImage = () =>
  // style.scssファイルを取得
  gulp.src("docs/img/**/*")
    .pipe(imgmin([
        imgminPNG({quality:  [.6, .8], speed: 1}),
        imgminJPG(),
        imgminGIF({
            interlaced: false,
            optimizationLevel: 3,
            colors:180
        }),
        imgminSVG()
        ]))
    .pipe(imgmin())//pngquantで暗くなる現象対応
    .pipe(gulp.dest("docs/img"));

// const watchSassFiles = () => watch("css/style.scss", compileSass);

// npx gulpというコマンドを実行した時、watchSassFilesが実行されるようにします
gulp.task(minifyImage);
// gulp.task(deploy);
// gulp.task('default', minifyImage);