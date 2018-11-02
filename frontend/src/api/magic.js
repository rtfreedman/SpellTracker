import axios from 'axios'

export default {
  getSlots(classes) {
    return axios
      .post('/api/magic/slots', { classes })
      .then(response => response.data)
  },

  getSpell(spell) {
    return axios
      .get(`/api/magic/spell/${spell}`)
      .then(response => response.data)
  },

  getSpells(spellClass, filter) {
    const body = {
      spellClass,
      filter,
    }
    return axios.get('/api/magic/spells', body).then(response => response.data)
  },

  getClasses() {
    return axios.get('/api/magic/classes').then(response => response.data)
  },
}
