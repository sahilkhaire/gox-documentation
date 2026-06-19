<script setup lang="ts">
defineProps<{
  pkg: string
  title: string
  node?: string
  importPath?: string
}>()
</script>

<template>
  <div class="symbol-header">
    <nav class="symbol-breadcrumb" aria-label="Breadcrumb">
      <a href="/packages/cond/">Packages</a>
      <span class="sep">/</span>
      <a :href="`/packages/${pkg}/`">{{ pkg }}</a>
      <span class="sep">/</span>
      <span class="current">{{ title }}</span>
    </nav>
    <div class="api-meta">
      <span v-if="node" class="api-badge node">Node: {{ node }}</span>
      <span v-if="importPath" class="api-badge import">{{ importPath }}</span>
    </div>
    <h1>{{ title }}</h1>
  </div>
</template>
