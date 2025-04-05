import { QueryClient } from '@tanstack/react-query'
import Swal from 'sweetalert2'

let displayedNetworkFailureError = false

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry(failureCount) {
        if (failureCount >= 3) {
          if (displayedNetworkFailureError === false) {
            displayedNetworkFailureError = true

            Swal.fire({
              icon: 'error',
              title: 'Erro de Conexão',
              text: 'A aplicação está demorando mais que o esperado para carregar, tente novamente em alguns minutos.',
              confirmButtonText: 'Ok',
              willClose: () => {
                displayedNetworkFailureError = false
              },
            })
          }

          return false
        }

        return true
      },
    },
  },
})
