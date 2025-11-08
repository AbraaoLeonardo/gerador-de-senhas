<template>
  <div class="app">
    <header class="header">
      <h1>Password Generator</h1>
    </header>

    <main class="content">
      <div class="form">
        <label for="size">Password Length:</label>
        <input
          id="size"
          type="number"
          v-model.number="size"
          min="1"
          max="30"
          placeholder="Enter size (max 30)"
        />

        <button @click="generatePassword">Generate</button>
      </div>

      <div class="result" v-if="password">
        <h3>Generated Password:</h3>
        <input type="text" :value="password" readonly />
      </div>
    </main>

    <footer class="footer">
      <p>© 2025 Password Generator App</p>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import './assets/style.css' // ✅ import external CSS file

const size = ref(16)
const password = ref('')

async function generatePassword() {
  try {
    const res = await fetch('http://localhost:8080/password/', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ size: size.value }),
    })

    if (!res.ok) throw new Error('Request failed')

    const data = await res.json()
    password.value = data.password
  } catch (err) {
    alert('Error generating password.')
    console.error(err)
  }
}
</script>
