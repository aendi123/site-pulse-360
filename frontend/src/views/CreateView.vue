<template>
  <form action="/api/website" method="post">
    <div class="m-3">
      <label for="name" class="form-label">Name</label>
      <input type="text" id="name" name="name" class="form-control" required>
    </div>
    <div class="m-3">
      <label for="url" class="form-label">URL</label>
      <input type="text" id="url" name="url" class="form-control" required>
    </div>
    <div class="m-3">
      <button type="submit" class="btn btn-primary">Submit</button>
    </div>
  </form>
</template>
