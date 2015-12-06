$(document).ready(function() {

	$('#form_trigger_id_sentence').submit(function(e) {
		e.preventDefault();
		$.ajax({
			type: "POST",
			url: "/main/post_form_sentence",
			data: $(this).serialize(),
			success: function(myData) {
				$.ajax({
					type: "GET",
					url: "/main/get_form_sentence",
					async: true,
					dataType: "json",
					success: function(myData) {
						var elem = document.getElementById('Sentence');
						elem.style.fontSize = "17px";
						elem.style.fontVariant = "bold";
						if (!myData.Sentence) {
							document.getElementsByName('form_name_sentence')[0].placeholder = 'please type anything...';
							elem.innerHTML = ""
						} else {
							elem.innerHTML = myData.Sentence;
						}
					}
				});
			}
		});
	});

	$('#button_trigger_id_reset').click(function(e) {
		e.preventDefault();
		$.ajax({
			type: "GET",
			url: "/main/reset",
			async: true,
			dataType: "json",
			success: function(myData) {
				var elem = document.getElementById('Sentence');
				elem.style.fontSize = "17px";
				elem.style.fontVariant = "bold";
				elem.innerHTML = "";
				document.getElementsByName('form_name_sentence')[0].placeholder = 'please type anything...';
			}
		});
	});

});