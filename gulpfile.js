const gulp = require('gulp');
const concat = require('gulp-concat');
const autoprefixer = require('gulp-autoprefixer');
const terser = require('gulp-terser');
const del = require('del');
const browserSync = require('browser-sync').create();
const plumber = require('gulp-plumber');
const notify = require("gulp-notify");
const fileinclude = require('gulp-file-include');
const replace = require('gulp-replace');
const htmlmin = require('gulp-htmlmin');
const cleanCSS = require('gulp-clean-css');
const dartSass =  require('sass');
const gulpSass = require('gulp-sass');
const sass = gulpSass(dartSass);
/*HTML*/
function html(){
	return gulp.src('./src/*.html')
	.pipe(plumber({errorHandler: notify.onError("Error: <%= error.message %>")}))
	.pipe(fileinclude({
		prefix: '@@',
		basepath: '@file'
	}))
	.pipe(gulp.dest('./build/'))
	.pipe(browserSync.stream());
}
function htmlMin(){
	return gulp.src('./build/*.html')
	.pipe(htmlmin({ collapseWhitespace: true }))
	.pipe(gulp.dest('./build/'));
}
/*JS*/
const jsFiles = [
	'./node_modules/slick-carousel/slick/slick.min.js',
	'./node_modules/magnific-popup/dist/jquery.magnific-popup.min.js',
	'./node_modules/bootstrap/dist/js/bootstrap.min.js',
	'./node_modules/air-datepicker/dist/js/datepicker.min.js',
	'./node_modules/tilt.js/dest/tilt.jquery.min.js',
	'./node_modules/jquery-countto/jquery.countTo.js',
	'./src/external/perfect-scrollbar/perfect-scrollbar.min.js',
	'./node_modules/imagesloaded/imagesloaded.pkgd.min.js',
	'./node_modules/lazysizes/lazysizes.min.js',
	'./node_modules/lazysizes/plugins/bgset/ls.bgset.min.js',
	'./src/external/panelmenu/panelmenu.js',
	'./src/external/form/jquery.form.js',
	'./src/external/form/jquery.validate.min.js',
	'./src/external/form/jquery.form-init.js',
	'./src/js/**/*.js'
];
function js() {
	return gulp.src(jsFiles, { allowEmpty: true })
	.pipe(concat('bundle.js'))
	.pipe(gulp.dest('./build/js'))
	.pipe(browserSync.stream());
}
function jsMin(){
	return gulp.src(jsFiles, { allowEmpty: true })
	.pipe(concat('bundle.js'))
	.pipe(terser({
		keep_fnames: true,
		mangle: false
	}))
	.pipe(gulp.dest('./build/js'))
	.pipe(browserSync.stream());
}
/*CSS*/
function css(){
	return gulp.src([
		'./src/scss/style.scss'
	])
	.pipe(plumber({errorHandler: notify.onError("Error: <%= error.message %>")}))
	.pipe(sass().on('error', sass.logError))
	.pipe(autoprefixer({
		overrideBrowserslist: ['last 1 versions'],
		cascade: false
	}))
	.pipe(gulp.dest('./build/css'))
	.pipe(gulp.dest('./src/css'))
	.pipe(browserSync.stream());
}
function cssMin(){
	return gulp.src([
		'./src/scss/style.scss'
	])
	.pipe(plumber({errorHandler: notify.onError("Error: <%= error.message %>")}))
	.pipe(sass().on('error', sass.logError))
	.pipe(autoprefixer({
		overrideBrowserslist: ['last 1 versions'],
		cascade: false
	}))
	.pipe(cleanCSS({level: 2}))
	.pipe(gulp.dest('./build/css'))
	.pipe(browserSync.stream());
}
/*Separate Css*/
function separateCss(done){
	gulp.src("./src/separate-include/**/*.scss")
	.pipe(sass().on('error', sass.logError))
	.pipe(autoprefixer({
		overrideBrowserslist: ['last 1 versions'],
		cascade: false
	}))
	.pipe(gulp.dest(function(file) {
		return file.base;
	}));
	fileTransfer2();
	done();
}
function fileTransfer2() {
	return gulp.src(['./src/separate-include/**/*'])
	.pipe(gulp.dest('./build/separate-include'))
	.pipe(browserSync.stream());
}
/*Remove Folder*/
function clean(){
	return del(['build/*'])
}
/*File Transfer*/
function fileTransfer() {
	return gulp.src(['./src/**/*', '!./src/js/**/*', '!./src/*.html'])
	.pipe(gulp.dest('./build/'))
	.pipe(browserSync.stream());
}
function watch() {
	gulp.watch('./src/scss/**/*.scss', css)
	gulp.watch('./src/separate-include/**/*.scss', separateCss)
	gulp.watch('./src/js/**/*.js', js)
	gulp.watch('./src/**/*.html', gulp.series(html)).on('change', browserSync.reload);
	browserSync.init({
		server: {
			baseDir: "./build/"
		}
	});
}
/*Task*/
gulp.task('css', css);
gulp.task('separateCss', separateCss);
gulp.task('html', html);
gulp.task('js', js);
gulp.task('del', clean);
gulp.task('fileTransfer', fileTransfer);
gulp.task('watch', watch);
gulp.task('build', gulp.series(clean, gulp.parallel(css,js,html,separateCss,fileTransfer)));
/*
	Task
	*Code minimization.
	**For website optimization
*/
gulp.task('htmlMin', htmlMin);
gulp.task('jsMin', jsMin);
gulp.task('cssMin', cssMin);
gulp.task('minAll', gulp.series(htmlMin,jsMin, cssMin));
