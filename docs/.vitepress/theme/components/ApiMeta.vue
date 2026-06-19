<script setup lang="ts">
defineProps<{
  pkg: string
  node?: string
  importPath?: string
}>()
</script>

<template>
  <div class="api-meta">
    <span v-if="node" class="api-badge node">Node: {{ node }}</span>
    <span v-if="importPath" class="api-badge import">{{ importPath }}</span>
    <span v-else-if="pkg" class="api-badge import">github.com/sahilkhaire/gox/{{ pkg }}</span>
  </div>
</template>
