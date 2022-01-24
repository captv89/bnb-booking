'use strict';

let attention = Prompt();

// Notify
function notify(type, msg) {
	notie.alert({
		type: type,
		text: msg,
		time: 3,
	});
}

// Modal Sweet Alert
function Prompt() {
	// Toast function
	const toast = function (c) {
		const { msg = '', icon = 'success', position = 'top-end' } = c;

		const Toast = Swal.mixin({
			toast: true,
			position: position,
			icon: icon,
			title: msg,
			showConfirmButton: false,
			timer: 3000,
			timerProgressBar: true,
			didOpen: (toast) => {
				toast.addEventListener('mouseenter', Swal.stopTimer);
				toast.addEventListener('mouseleave', Swal.resumeTimer);
			},
		});
		Toast.fire({});
	};
	// Popup function
	const popmsg = function (c) {
		const {
			msg = '',
			icon = 'success',
			title = '',
			btnText = 'OK',
			type = 'text',
		} = c;
		if (type === 'text') {
			Swal.fire({
				icon: icon,
				title: title,
				text: msg,
				confirmButtonText: btnText,
			});
		} else {
			Swal.fire({
				icon: icon,
				title: title,
				html: msg,
				confirmButtonText: btnText,
			});
		}
	};

	// Pop form
	const popform = async function (c) {
		const { msg = '', title = '' } = c;

		const { value: formValues } = await Swal.fire({
			title: title,
			html: msg,
			focusConfirm: false,
			showCancelButton: true,
			willOpen: () => {
                // Get Date Picker for Swal
				const elem = document.getElementById('reservation-dates');
				const rangepicker = new DateRangePicker(elem, {
					format: 'dd-mm-yyyy',
					todayBtn: 'true',
					clearBtn: 'true',
					buttonClass: 'btn',
					calendarWeeks: 'true',
					minDate: 'today',
					autohide: 'true',
				});
			},
			preConfirm: () => {
				return [
					document.getElementById('check-in').value,
					document.getElementById('check-out').value,
				];
			},
			didOpen: () => {
				document.getElementById('check-in').removeAttribute('disabled');
				document
					.getElementById('check-out')
					.removeAttribute('disabled');
			},
		});

		if (formValues) {
			Swal.fire(JSON.stringify(formValues));
		}
	};

	return {
		toast: toast,
		popmsg: popmsg,
		popform: popform,
	};
}