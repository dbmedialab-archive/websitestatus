var gulp = require('gulp'),
browserify = require('browserify'),
babelify = require('babelify'),
sourcemaps = require('gulp-sourcemaps'),
notify = require('gulp-notify'),
buffer = require('gulp-buffer'),
uglify = require('gulp-uglify'),
concat = require('gulp-concat'),
autoprefixer = require('gulp-autoprefixer'),
source = require('vinyl-source-stream'),
plumber = require('gulp-plumber');

gulp.task('js', function() {
    browserify('./src/assets/js/app.js')
        .transform(babelify, {presets: ["es2015", "react"]})
        .bundle().on('error', notify.onError( {
            title: 'JS Error',
            message: "<%= error.message %>"
        }))
        .pipe(source('app.min.js'))
        .pipe(buffer())
        .pipe(sourcemaps.init({loadMaps: true}))
        .pipe(uglify())
        .pipe(sourcemaps.write('.'))
        .pipe(gulp.dest('./static/assets/js/'));
});

gulp.task('build', ['js']);

gulp.task('watch', function() {
    gulp.watch('./src/assets/js/app.js', ['js']);
});

gulp.task('default', ['js', 'watch']);