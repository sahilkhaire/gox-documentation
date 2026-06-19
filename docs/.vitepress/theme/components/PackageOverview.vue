<script setup lang="ts">
defineProps<{
  name: string
  node?: string
  importPath: string
  subtitle?: string
  symbolCount?: number
  hasCookbook?: boolean
  migrationLink?: string
  narrative?: string
}>()
</script>

<template>
  <div class="package-overview">
    <div class="package-hero">
      <div class="api-meta">
        <span v-if="node" class="api-badge node">Node: {{ node }}</span>
        <span v-if="hasCookbook" class="api-badge cookbook">Cookbook</span>
        <span v-if="migrationLink" class="api-badge migration">Migration guide</span>
      </div>
      <h1>{{ name }}</h1>
      <p v-if="subtitle" class="subtitle">{{ subtitle }}</p>
      <div class="import-line">import "{{ importPath }}"</div>
    </div>

    <div v-if="symbolCount" class="package-stats">
      <div class="package-stat">
        <div class="value">{{ symbolCount }}</div>
        <div class="label">APIs documented</div>
      </div>
      <div v-if="hasCookbook" class="package-stat">
        <div class="value">📖</div>
        <div class="label"><a :href="`/packages/${name}/cookbook`">Cookbook</a></div>
      </div>
      <div v-if="migrationLink" class="package-stat">
        <div class="value">→</div>
        <div class="label"><a :href="migrationLink">From Node.js</a></div>
      </div>
    </div>

    <div v-if="narrative" class="package-narrative">
      <h2>Why this package</h2>
      <p>{{ narrative }}</p>
    </div>
  </div>
</template>
