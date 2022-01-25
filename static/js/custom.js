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
				if (c.willOpen !== undefined) {
					c.willOpen();
				}
			},
			preConfirm: () => {
				if (c.preConfirm !== undefined) {
					c.preConfirm();
				}
			},
			didOpen: () => {
				if (c.didOpen !== undefined) {
					c.didOpen();
				}
			},
		});

		if (formValues) {
			if (formValues !== Swal.DismissReason.cancel) {
				if (formValues.length > 0) {
					if (c.callback !== undefined) {
						c.callback(formValues);
					}
				} else {
					c.callback(false);
				}
			} else {
				c.callback(false);
			}
		}
	};

	return {
		toast: toast,
		popmsg: popmsg,
		popform: popform,
	};
}
