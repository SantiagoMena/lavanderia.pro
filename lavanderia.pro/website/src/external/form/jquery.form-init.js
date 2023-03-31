jQuery(function($) {
	function successModal(){
		if($('body').hasClass('modal-open')){
			var objModal = $('.modal');
			if(!objModal.hasClass('show')) return false;
			objModal.find('.tt-modal-message').addClass('tt-active');
			setTimeout(function(){
				objModal.find('.close').trigger('click');
				objModal.find('.tt-modal-message').removeClass('tt-active');
			}, 1600);
		};
	};
	function successForm(obj){
		var objMessage = $(obj).find('.tt-modal-message');
		objMessage.addClass('tt-active');
		setTimeout(function(){
			objMessage.removeClass('tt-active');
		}, 3600);
	};
	var formContactForm = $('#contact-form');
	if (formContactForm.length){
		formContactForm.validate({
			rules: {
				name: {
					required: true,
					minlength: 2
				},
				email: {
					required: true,
					email: true
				},
				message: {
					required: true,
				}
			},
			messages: {
				name: {
					required: "Please enter your name",
					minlength: "Your name must consist of at least 2 characters"
				},
				email: {
					required: "Please enter your email"
				},
				message: {
					required: "Please enter your message"
				}
			},
			submitHandler: function(form) {
				$(form).ajaxSubmit({
					type:"POST",
					data: $(form).serialize(),
					url:"external/form/contact-form.php",
					success: function() {
						$('#success').fadeIn();
						formContactForm.each(function(){
							this.reset();
						});
						successForm(formContactForm);
					},
					error: function() {
						formContactForm.fadeTo( "slow", 1, function() {
							$('#error').fadeIn();
						});
					}
				});
			}
		});
	};
	var subscribeform = $('#subscribeform');
	if (subscribeform.length){
		subscribeform.validate({
			rules: {
				email: {
					required: true,
					email: true
				}
			},
			submitHandler: function(form) {
				$(form).ajaxSubmit({
					type:"POST",
					data: $(form).serialize(),
					url:"external/form/newsletter-form.php",
					success: function() {
						$('#success').fadeIn();
						subscribeform.each(function(){this.reset();});
						successForm(subscribeform);
					},
					error: function() {
						subscribeform.fadeTo( "slow", 1, function() {
							$('#error').fadeIn();
						});
					}
				});
			}
		});
	};
	var faqForm = $('#faq-form');
	if (faqForm.length){
		faqForm.validate({
			rules: {
				name: {
					required: true,
					minlength: 2
				},
				email: {
					required: true,
					email: true
				},
				message: {
					required: true,
				}
			},
			messages: {
				name: {
					required: "Please enter your name",
					minlength: "Your name must consist of at least 2 characters"
				},
				email: {
					required: "Please enter your email"
				},
				message: {
					required: "Please enter your message"
				}
			},
			submitHandler: function(form) {
				$(form).ajaxSubmit({
					type:"POST",
					data: $(form).serialize(),
					url:"external/form/faq-form.php",
					success: function() {
						$('#success').fadeIn();
						faqForm.each(function(){
							this.reset();
						});
						successForm(subscribeform);
					},
					error: function() {
						faqForm.fadeTo( "slow", 1, function() {
							$('#error').fadeIn();
						});
					}
				});
			}
		});
	};
	//jsFormRequestQuote
	var formModal = $('#jsFormRequestQuote');
	if (formModal.length){
		formModal.validate({
			rules: {
				name: {
					required: true,
					minlength: 2
				},
				email: {
					required: true,
					email: true
				},
				message: {
					required: true,
				}
			},
			messages: {
				name: {
					required: "Please enter your name",
					minlength: "Your name must consist of at least 2 characters"
				},
				email: {
					required: "Please enter your email"
				},
				message: {
					required: "Please enter your message"
				}
			},
			submitHandler: function(form) {
				$(form).ajaxSubmit({
					type:"POST",
					data: $(form).serialize(),
					url:"external/form/modal-schedule-pickup.php",
					success: function(){
						$('#success').fadeIn();
						formModal.each(function(){this.reset();});
						successModal(formModal);
					},
					error: function() {
						formModal.fadeTo( "slow", 1, function(){
							$('#error').fadeIn();
						});
					}
				});
			}
		});
	};
	//blog-item.html
	var blogForm = $('#feedbackComment');
	if (blogForm.length){
		blogForm.validate({
			rules: {
				name: {
					required: true,
					minlength: 2
				},
				email: {
					required: true,
					email: true
				},
				message: {
					required: true,
				}
			},
			messages: {
				name: {
					required: "Please enter your name",
					minlength: "Your name must consist of at least 2 characters"
				},
				email: {
					required: "Please enter your email"
				},
				message: {
					required: "Please enter your message"
				}
			},
			submitHandler: function(form) {
				$(form).ajaxSubmit({
					type:"POST",
					data: $(form).serialize(),
					url:"external/form/comment-form.php",
					success: function() {
						$('#success').fadeIn();
						blogForm.each(function(){
							this.reset();
						});
						successForm(subscribeform);
					},
					error: function() {
						blogForm.fadeTo( "slow", 1, function() {
							$('#error').fadeIn();
						});
					}
				});
			}
		});
	};
});