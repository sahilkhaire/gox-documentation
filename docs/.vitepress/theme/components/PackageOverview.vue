<script setup lang="ts">
defineProps<{
  name: string
  node?: string
  importPath: string
  subtitle?: string
  symbolCount?: number
}>()
</script>

<template>
  <div class="package-overview">
    <div class="package-hero">
      <div class="api-meta">
        <span v-if="node" class="api-badge node">Node: {{ node }}</span>
        <span v-if="symbolCount" class="api-badge import">{{ symbolCount }} APIs</span>
      </div>
      <h1>{{ name }}</h1>
      <p v-if="subtitle" class="subtitle">{{ subtitle }}</p>
      <div class="import-line">import "{{ importPath }}"</div>
    </div>
  </div>
</template>
