var gulp = require('gulp'),
browserify = require('browserify'),
vueify = require('vueify'),
source = require('vinyl-source-stream'),
notify = require('gulp-notify'),
buffer = require('gulp-buffer'),
plumber = require('gulp-plumber');
sourcemaps = require('gulp-sourcemaps'),
uglify = require('gulp-uglify'),

gulp.task('js', function() {
    browserify('./src/assets/js/app.js')
        .transform(vueify)
        .bundle().on('error', notify.onError( {
            title: 'JS Error',
            message: "<%= error.message %>"
        }))
        .pipe(source('bundle.js'))
        .pipe(buffer())
        .pipe(sourcemaps.init({loadMaps: true}))
        .pipe(uglify())
        .pipe(sourcemaps.write('.'))
        .pipe(gulp.dest('./static/assets/js/'));
});

gulp.task('style', function () {

});

gulp.task('build', ['js']);

gulp.task('watch', function() {
    gulp.watch('./src/assets/js/components/*.vue', ['js']);
    gulp.watch('./src/assets/js/app.js', ['js']);
});

gulp.task('default', ['js']);
