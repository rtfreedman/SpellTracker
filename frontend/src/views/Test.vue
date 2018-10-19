<template>
  <v-container>

    <v-tabs v-model="activeTab"
            color="cyan"
            dark
            slider-color="yellow">
      <v-tab v-for="n in tabs.length"
             :key="'tab' + n"
             ripple>
        <span>{{ tabs[n - 1].name }}</span>
        <v-icon class="ml-3"
                @click.stop="remove(n)">close</v-icon>
      </v-tab>
      <v-tab-item v-for="n in tabs.length"
                  :key="'item' + n">
        <v-card flat>
          <v-card-text>{{ tabs[n - 1].content }}</v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs>

    <div class="text-xs-center mt-3">
      <v-btn @click="next">next tab</v-btn>
      <v-btn @click="add">add tab</v-btn>
    </div>

    <pre>activeTab: {{ activeTab }}</pre>
    <pre>tabs: {{ tabs }}</pre>

    <v-divider class="my-3"></v-divider>

    <h2 class="mb-2">Better?</h2>

    <v-tabs v-model="activeTab"
            color="cyan"
            dark
            slider-color="yellow">
      <v-tab v-for="n in tabs.length"
             :key="'tab' + n"
             ripple>
        <span>{{ tabs[n - 1].name }}</span>
      </v-tab>
      <v-tab-item v-for="n in tabs.length"
                  :key="'item' + n">
        <v-card flat>
          <v-layout justify-end>
            <v-icon class="ma-3"
                  @click.stop="remove(n)">close</v-icon>
          </v-layout>
          <v-card-text>{{ tabs[n - 1].content }}</v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs>

    <div class="text-xs-center mt-3">
      <v-btn @click="next">next tab</v-btn>
      <v-btn @click="add">add tab</v-btn>
    </div>

    <pre>activeTab: {{ activeTab }}</pre>
    <pre>tabs: {{ tabs }}</pre>

  </v-container>
</template>

<script>
export default {
  data() {
    return {
      tabIndex: 4,
      tabs: [
        { name: 'tab_1', content: 'TAB 1 CONTENT' },
        { name: 'tab_2', content: 'TAB 2 CONTENT' },
        { name: 'tab_3', content: 'TAB 3 CONTENT' },
      ],
      activeTab: 0,
    }
  },
  methods: {
    next() {
      const current = parseInt(this.activeTab)
      this.activeTab = current < this.tabs.length - 1 ? current + 1 : 0
    },

    add() {
      this.tabs.push({
        name: `tab_${this.tabIndex}`,
        content: `TAB ${this.tabIndex} CONTENT`,
      })
      this.tabIndex += 1
    },

    remove(n) {
      this.tabs.splice(n - 1, 1)
      
      // this isn't necessary in the second implementation
      if (n === 1) {
        this.activeTab = 0
      }
    },
  },
}
</script>

<style scoped>
</style>
