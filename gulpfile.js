'use strict';

// $ npm install gulp gulp-sass gulp-concat --save-dev

var gulp = require('gulp');
var sass = require('gulp-sass');
var concat = require('gulp-concat');
var inject = require('gulp-inject');

const sassPath = './public/assets/sass/**/*.scss';
const cssPath = './public/assets';
const cssFile = 'style.css';

// $ gulp sass

gulp.task('sass', function () {
  return gulp.src(sassPath)
    .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
    .pipe(concat(cssFile))
    .pipe(gulp.dest(cssPath));
	});

// $ gulp sass:watch

gulp.task('sass:watch', function () {
  gulp.watch('sassPath', ['sass']);
	});


// $ gulp injectjs

gulp.task('injectjs', function () {
	var target = gulp.src('./public/index.html');
	var sources = gulp.src(['./public/assets/js/**/*.js'], {read: false});

	return target.pipe(inject(sources))
		.pipe(gulp.dest('./public'));
	});