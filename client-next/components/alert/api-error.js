import Swal from 'sweetalert2'

const ApiErrorAlert = (title, text) => {
  Swal.fire({
    title,
    text,
    icon: 'error',
    focusConfirm: false,
  })
}

export default ApiErrorAlert
