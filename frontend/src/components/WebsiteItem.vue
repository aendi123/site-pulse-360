<script lang="ts">
export default {
  props: ['id', 'name', 'url', 'up']
}
</script>

<template>
  <RouterLink :to="'/website/' + id" class="text-decoration-none text-body">
    <div class="website-item-container shadow p-3 mb-3 bg-body-tertiary rounded bg-light-subtle">
      <span class="website-name fs-4">{{ name }}</span>
      <span class="website-url"><a :href="url" class="text-secondary" @click.stop>{{ url }}</a></span>
      <span class="website-status">
        <img v-if="up" src="../assets/up.png" alt="Website is up" height="80%">
        <img v-else src="../assets/down.png" alt="Website is down" height="80%">
      </span>    
    </div>
  </RouterLink>
</template>

<style scoped>
.website-item-container {
  display: grid;
  grid-template-columns: 4fr 1fr;
  grid-template-rows: 80px 20px;
  grid-template-areas: "name status"
  "url  status";
}

.website-name {
  grid-area: name;
}

.website-url {
  grid-area: url;
}

.website-status {
  grid-area: status;
  display: flex;
  justify-content: right;
  align-items: center;
}
</style>