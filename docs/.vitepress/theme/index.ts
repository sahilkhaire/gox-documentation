import DefaultTheme from 'vitepress/theme'
import './custom.css'
import ApiMeta from './components/ApiMeta.vue'
import PackageHero from './components/PackageHero.vue'
import HomeStats from './components/HomeStats.vue'
import TierGrid from './components/TierGrid.vue'

export default {
  extends: DefaultTheme,
  enhanceApp({ app }) {
    app.component('ApiMeta', ApiMeta)
    app.component('PackageHero', PackageHero)
    app.component('HomeStats', HomeStats)
    app.component('TierGrid', TierGrid)
  }
}
