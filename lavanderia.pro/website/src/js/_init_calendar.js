(function($){
	$.fn.datepicker.language['en'] = {
		days: ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'],
		daysShort: ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'],
		daysMin: ['Su', 'Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa'],
		months: ['January','February','March','April','May','June', 'July','August','September','October','November','December'],
		monthsShort: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
		today: 'Today',
		clear: 'Clear',
		dateFormat: 'mm/dd/yyyy',
		timeFormat: 'hh:ii aa',
		firstDay: 0
	};
	var modalObj = $('#modalMRequestQuote');
	var dp = modalObj.find('.js_datepicker-1').datepicker({
		language: 'en',
	}).data('datepicker');
	var dp = modalObj.find('.js_datepicker-2').datepicker({
		language: 'en',
	}).data('datepicker');
})(jQuery);



