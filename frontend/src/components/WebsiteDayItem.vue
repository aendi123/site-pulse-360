<script lang="ts">
export default {
  props: ['date', 'uptime', 'avgResponseTime']
}
</script>

<template>
<tr :class="{
    'table-success': uptime === 100,
    'table-warning': uptime < 100 && uptime >= 99,
    'table-danger': uptime < 99
  }">
  <td>{{ date }}</td>
  <td>{{ uptime }}%</td>
  <td>{{ avgResponseTime }}ms</td>
</tr>
</template>