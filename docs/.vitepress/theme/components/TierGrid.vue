<template>
  <div class="tier-grid">
    <a class="tier-card" href="/packages/cond/">
      <div class="tier-title">Utilities</div>
      <div class="tier-packages">cond, slice, maputil, str, set, buffer, async, path, url</div>
    </a>
    <a class="tier-card" href="/packages/env/">
      <div class="tier-title">I/O &amp; OS</div>
      <div class="tier-packages">env, fs, err, exec, osutil, stream, compress, csv, archive</div>
    </a>
    <a class="tier-card" href="/packages/http/">
      <div class="tier-title">Web</div>
      <div class="tier-packages">http, client, log, ws — Express-style routes &amp; middleware</div>
    </a>
    <a class="tier-card" href="/packages/db/">
      <div class="tier-title">Data</div>
      <div class="tier-packages">json, validate, db, redis, mongo</div>
    </a>
    <a class="tier-card" href="/packages/jwt/">
      <div class="tier-title">Security</div>
      <div class="tier-packages">crypto, jwt, id, auth</div>
    </a>
    <a class="tier-card" href="/packages/queue/">
      <div class="tier-title">Infrastructure</div>
      <div class="tier-packages">cache, event, cron, queue, mail, time, semver</div>
    </a>
  </div>
</template>

<style scoped>
a.tier-card {
  text-decoration: none;
  color: inherit;
  display: block;
}
</style>
