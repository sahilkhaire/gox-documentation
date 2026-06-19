import { defineConfig } from 'vitepress'
import { generatedSidebar } from './sidebar.generated'

export default defineConfig({
  title: 'gox',
  description: 'Node-friendly Go toolkit — premium documentation for every package and API',
  base: '/',
  lang: 'en-US',
  sitemap: {
    hostname: 'https://gox.varcore.dev'
  },
  head: [
    ['link', { rel: 'icon', href: '/favicon.svg' }],
    ['meta', { name: 'theme-color', content: '#00add8' }],
    ['link', { rel: 'canonical', href: 'https://gox.varcore.dev' }]
  ],
  ignoreDeadLinks: true,
  outline: { level: [2, 3] },
  themeConfig: {
    logo: { src: '/logo.svg', width: 28, height: 28 },
    siteTitle: 'gox',
    nav: [
      { text: 'Guide', link: '/guide/getting-started' },
      { text: 'Cookbooks', link: '/guide/cookbooks' },
      { text: 'Cheat Sheet', link: '/reference/cheatsheet' },
      { text: 'Packages', link: '/packages/cond/' },
      { text: 'Migration', link: '/migration/express' },
      {
        text: 'GitHub',
        link: 'https://github.com/sahilkhaire/gox'
      }
    ],
    sidebar: {
      '/guide/': [
        {
          text: 'Guide',
          items: [
            { text: 'Getting Started', link: '/guide/getting-started' },
            { text: 'Cookbooks', link: '/guide/cookbooks' },
            { text: 'Architecture', link: '/guide/architecture' },
            { text: 'HTTP Guide', link: '/guide/http' }
          ]
        }
      ],
      '/reference/': [
        {
          text: 'Reference',
          items: [{ text: 'Cheat Sheet', link: '/reference/cheatsheet' }]
        }
      ],
      '/migration/': [
        {
          text: 'Migration Guides',
          items: [
            { text: 'Express', link: '/migration/express' },
            { text: 'axios', link: '/migration/axios' },
            { text: 'Knex', link: '/migration/knex' },
            { text: 'Passport', link: '/migration/passport' },
            { text: 'WebSocket', link: '/migration/ws' },
            { text: 'GraphQL', link: '/migration/graphql' },
            { text: 'gRPC', link: '/migration/grpc' },
            { text: 'Kafka', link: '/migration/kafka' },
            { text: 'OpenTelemetry', link: '/migration/opentelemetry' },
            { text: 'Testing (Jest)', link: '/migration/testing' }
          ]
        }
      ],
      ...generatedSidebar
    },
    search: { provider: 'local' },
    socialLinks: [
      { icon: 'github', link: 'https://github.com/sahilkhaire/gox' }
    ],
    footer: {
      message: 'MIT Licensed · Built for Node.js developers moving to Go',
      copyright: 'Copyright © gox contributors'
    },
    docFooter: { prev: 'Previous', next: 'Next' }
  }
})
