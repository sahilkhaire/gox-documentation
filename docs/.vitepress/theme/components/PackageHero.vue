<script setup lang="ts">
defineProps<{
  name: string
  node?: string
  importPath: string
  subtitle?: string
}>()
</script>

<template>
  <div class="package-hero">
    <ApiMeta :pkg="name" :node="node" :import-path="importPath" />
    <h1>{{ name }}</h1>
    <p v-if="subtitle" class="subtitle">{{ subtitle }}</p>
    <div class="import-line">import "{{ importPath }}"</div>
  </div>
</template>
