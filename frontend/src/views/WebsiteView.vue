<script setup lang="ts">
import WebsiteDayItem from '../components/WebsiteDayItem.vue';
</script>

<script lang="ts">
import { useRoute } from 'vue-router';

export default {
    data() {
        return {
            websiteid: Number(useRoute().params.id) as number,
            websitename: '' as string,
            dayoverview: {} as Record<string, { upPercentage: number, avgResponseTime: number }>
        };
    },
    methods: {
        fetchWebsiteName() {
            fetch(`/api/`)
            .then(response => response.json())
            .then(response => {
                this.websitename = response.find((website: any) => website.ID === this.websiteid).Name;
            });
        },
        
        fetchWebsiteData() {
            fetch(`/api/website/${this.websiteid}/state?limit=43200`)
            .then(response => response.json())
            .then(data => {
                const groupedByDay = data.reduce((acc: Record<string, any[]>, curr: any) => {
                    const date = new Date(curr.request_date);
                    const day = `${date.getDate()}.${date.getMonth()}.${date.getFullYear()}`;
                    
                    if (!acc[day]) {
                        acc[day] = [];
                    }
                    
                    acc[day].push(curr);
                    return acc;
                }, {});
                
                for (const day in groupedByDay) {
                    const dayData = groupedByDay[day];
                    const upCount = dayData.filter((item: any) => item.up).length;
                    const upPercentage = Number(((upCount / dayData.length) * 100).toFixed(2));
                    const avgResponseTime = Number((dayData.reduce((acc: number, curr: any) => acc + curr.response_time_ms, 0) / dayData.length).toFixed(2));
                    
                    this.dayoverview[day] = {
                        upPercentage,
                        avgResponseTime
                    };
                }
            });
        },
        
        async deletewebsite() {
            await fetch(`/api/website/${this.websiteid}`, { method: 'DELETE' });
            window.location.href = '/';
        }
    },
    created() {
        this.fetchWebsiteName();
        this.fetchWebsiteData();
    }
};
</script>

<template>
    <div class="website-name-container">
        <h1>{{ websitename }}</h1>
        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-trash3-fill" viewBox="0 0 16 16" @click="deletewebsite">
            <path d="M11 1.5v1h3.5a.5.5 0 0 1 0 1h-.538l-.853 10.66A2 2 0 0 1 11.115 16h-6.23a2 2 0 0 1-1.994-1.84L2.038 3.5H1.5a.5.5 0 0 1 0-1H5v-1A1.5 1.5 0 0 1 6.5 0h3A1.5 1.5 0 0 1 11 1.5m-5 0v1h4v-1a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5M4.5 5.029l.5 8.5a.5.5 0 1 0 .998-.06l-.5-8.5a.5.5 0 1 0-.998.06m6.53-.528a.5.5 0 0 0-.528.47l-.5 8.5a.5.5 0 0 0 .998.058l.5-8.5a.5.5 0 0 0-.47-.528M8 4.5a.5.5 0 0 0-.5.5v8.5a.5.5 0 0 0 1 0V5a.5.5 0 0 0-.5-.5"/>
        </svg>
    </div>
    <h2>History of the last 30 days</h2>
    <table class="table table-hover">
        <thead>
            <tr>
                <th>Date</th>
                <th>Uptime</th>
                <th>Avg response time</th>
            </tr>
        </thead>
        <tbody>
            <WebsiteDayItem
            v-for="(value, key) in dayoverview"
            :key="key"
            :date="key"
            :uptime="value.upPercentage"
            :avgResponseTime="value.avgResponseTime"
            />
        </tbody>
    </table>
</template>

<style scoped>
svg:hover {
    cursor: pointer;
}

.website-name-container {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
}

h2 {
    font-size: 20px;
    margin-bottom: 20px;
}
</style>