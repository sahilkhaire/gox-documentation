<script setup lang="ts">
import { onMounted, ref } from 'vue'

const props = defineProps<{
  placeholder?: string
}>()

const query = ref('')

onMounted(() => {
  const input = document.querySelector<HTMLInputElement>('.symbol-filter-input')
  if (!input) return
  input.addEventListener('input', () => {
    query.value = input.value.toLowerCase()
    document.querySelectorAll('.symbol-table tbody tr').forEach((row) => {
      const text = (row as HTMLElement).innerText.toLowerCase()
      ;(row as HTMLElement).style.display = text.includes(query.value) ? '' : 'none'
    })
  })
})
</script>

<template>
  <div class="symbol-filter">
    <input
      class="symbol-filter-input"
      type="search"
      :placeholder="placeholder || 'Filter symbols…'"
      autocomplete="off"
    />
  </div>
</template>
