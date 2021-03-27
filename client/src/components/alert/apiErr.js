import Swal from 'sweetalert2'
import './apiErr.scss'

const ApiErrAlert = (title, text) => {
  Swal.fire({
    title,
    text,
    icon: 'error',
    focusConfirm: false,
  })
}

export default ApiErrAlert
