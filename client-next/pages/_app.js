import { appWithTranslation } from 'next-i18next'
import { StoreProvider } from 'store'
import 'styles/globals.css'

function MyApp({ Component, pageProps }) {
  return (
    <StoreProvider>
      <Component {...pageProps} />
    </StoreProvider>
  )
}

export default appWithTranslation(MyApp)
