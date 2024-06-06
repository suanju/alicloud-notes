import { defineConfig ,presetUno,presetIcons} from 'unocss'

export default defineConfig({
  // ...UnoCSS选项
  presets: [
    presetUno(),
    presetIcons({
      collections: {
        systemUicons: () => import('@iconify-json/system-uicons/icons.json').then((i) => i.default)
      }
    })
  ]
})
