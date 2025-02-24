<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { ref } from 'vue';
import WebsiteItem from '../components/WebsiteItem.vue';

const websites : any = ref(null);

fetch('/api')
.then(response => response.json())
.then(response => {
    websites.value = response;
    for (const website of websites.value) {
        fetch(`/api/website/${website.ID}/state?limit=1`)
        .then(response => response.json())
        .then(response => website.up = response[0].up);
    }
});
</script>

<template>
    <WebsiteItem
    v-for="website in websites"
    :key="website.ID"
    :id="website.ID"
    :name="website.Name"
    :url="website.URL"
    :up="website.up"
    />
    <RouterLink to="/create" class="text-reset text-decoration-none fs-4">
        <div class="shadow p-3 bg-body-tertiary rounded bg-light-subtle text-center">+</div>
    </RouterLink>
    
    
</template>