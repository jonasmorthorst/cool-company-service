<template>
  <section v-if="company != null" class="section">
    <div class="container">
      <div v-if="errors.length" class="notification is-warning">
        <b>Please correct the following error(s):</b>
        <ul>
          <li v-for="(error, index) in errors" :key="index">
            {{ error }}
          </li>
        </ul>
      </div>

      <h1 v-if="!isNew" class="title">
        {{ company.name }}
      </h1>
      <h1 v-else class="title">
        Create company
      </h1>

      <form @submit="saveClicked">
        <div class="buttons">
          <button v-if="!isEdit" type="button" class="button is-primary is-outlined" @click="editClicked">
            <span class="icon is-small">
              <i class="fas fa-edit" />
            </span>
            <span>Edit</span>
          </button>

          <button v-if="isEdit && !isNew" type="submit" class="button is-primary">
            <span class="icon is-small">
              <i class="fas fa-check" />
            </span>
            <span>Save</span>
          </button>

          <button v-if="!isNew" type="button" class="button is-danger is-outlined" @click="deleteClicked">
            <span class="icon is-small">
              <i class="fas fa-times" />
            </span>
            <span>Delete</span>
          </button>
        </div>

        <div class="columns">
          <div class="column">
            <div class="field">
              <label class="label">Name</label>
              <div class="control">
                <span v-if="!isEdit">{{ company.name }}</span>
                <input
                  v-if="isEdit"
                  v-model="companyCopy.name"
                  class="input"
                  type="text"
                  placeholder="Name"
                  required
                >
              </div>
            </div>

            <div class="field">
              <label class="label">Company ID</label>
              <div class="control">
                <span v-if="!isEdit">{{ company.companyID }}</span>
                <input
                  v-if="isEdit"
                  v-model="companyCopy.companyID"
                  class="input"
                  type="text"
                  placeholder="Company ID"
                  required
                >
              </div>
            </div>

            <div class="field">
              <label class="label">E-mail</label>
              <div class="control">
                <span v-if="!isEdit">{{ company.email }}</span>
                <span v-if="!isEdit && company.email == ''"><i>No e-mail</i></span>
                <input v-if="isEdit" v-model="companyCopy.email" class="input" type="email" placeholder="E-mail">
              </div>
            </div>

            <div class="field">
              <label class="label">Phone</label>
              <div class="control">
                <span v-if="!isEdit">{{ company.phone }}</span>
                <span v-if="!isEdit && company.phone == ''"><i>No phone</i></span>
                <input v-if="isEdit" v-model="companyCopy.phone" class="input" type="text" placeholder="Phone">
              </div>
            </div>
          </div>
          <div class="column">
            <div class="field">
              <label class="label">Country</label>
              <div class="control">
                <span v-if="!isEdit">{{ company.country }}</span>
                <input
                  v-if="isEdit"
                  v-model="companyCopy.country"
                  class="input"
                  type="text"
                  placeholder="Country"
                  required
                >
              </div>
            </div>

            <div class="field">
              <label class="label">Address</label>
              <div class="control">
                <span v-if="!isEdit">{{ company.address }}</span>
                <input
                  v-if="isEdit"
                  v-model="companyCopy.address"
                  class="input"
                  type="text"
                  placeholder="Address"
                  required
                >
              </div>
            </div>

            <div class="field">
              <label class="label">City</label>
              <div class="control">
                <span v-if="!isEdit">{{ company.city }}</span>
                <input
                  v-if="isEdit"
                  v-model="companyCopy.city"
                  class="input"
                  type="text"
                  placeholder="City"
                  required
                >
              </div>
            </div>

            <div class="field">
              <label class="label">ZIP</label>
              <div class="control">
                <span v-if="!isEdit">{{ company.zipCode }}</span>
                <input
                  v-if="isEdit"
                  v-model="companyCopy.zipCode"
                  class="input"
                  type="text"
                  placeholder="ZIP"
                  required
                >
              </div>
            </div>
          </div>
        </div>

        <button v-if="isNew" type="submit" class="button is-primary">
          <span class="icon is-small">
            <i class="fas fa-check" />
          </span>
          <span>Create</span>
        </button>
      </form>

      <hr>

      <div v-if="!isNew" class="columns">
        <div class="column is-half">
          <h1 class="title is-5">
            Owners
          </h1>

          <div class="table-container">
            <table class="table is-fullwidth is-hoverable is-striped">
              <tr>
                <th>Name</th>
                <th>Title</th>
                <th />
              </tr>
              <tr v-for="(owner, index) in company.owners" :key="index">
                <td>{{ owner.name }}</td>
                <td>{{ owner.title }}</td>
                <td><a @click="removeOwner(owner)">Remove</a></td>
              </tr>
              <tr v-if="!company.owners.length">
                <td colspan="3"><i>No owners</i></td>
              </tr>
            </table>
          </div>
        </div>
        <div class="column">
          <form @submit="addOwner">
            <h1 class="title is-5">
              Add owner
            </h1>
            <div class="field">
              <label class="label">Name</label>
              <div class="control">
                <input v-model="newOwner.name" class="input" type="text" placeholder="Name" required>
              </div>
            </div>

            <div class="field">
              <label class="label">Title</label>
              <div class="control">
                <input v-model="newOwner.title" class="input" type="text" placeholder="Title" required>
              </div>
            </div>
            <div class="control">
              <button class="button is-primary" type="submit">
                <span class="icon is-small">
                  <i class="fas fa-plus" />
                </span>
                <span>Add</span>
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>

<script>

export default {
  data () {
    return {
      company: null,
      companyCopy: {},
      isEdit: false,
      isNew: false,
      errors: [],
      newOwner: {}
    }
  },
  created () {
    // If new
    if (this.$route.params.id === '0') {
      this.isNew = true
      this.isEdit = true

      this.company = {
        name: '',
        companyID: '',
        email: '',
        phone: '',
        address: '',
        zipCode: '',
        city: '',
        country: '',
        owners: []
      }

      this.companyCopy = Object.assign({}, this.company)
    } else {
      this.getCompany()
    }
  },
  methods: {
    async getCompany () {
      try {
        const data = await this.$axios.$get('/companies/' + this.$route.params.id)

        this.company = data.company
      } catch (e) {
        this.$router.replace({ path: '/' })

        this.$toast.open({
          message: e,
          position: 'top-right',
          type: 'error'
        })
      }
    },
    editClicked (e) {
      e.preventDefault()
      this.companyCopy = Object.assign({}, this.company)
      this.isEdit = true
    },
    async saveClicked (e) {
      e.preventDefault()

      // Validate
      this.errors = []

      if (!this.companyCopy.companyID) { this.errors.push('Company ID required.') }
      if (!this.companyCopy.name) { this.errors.push('Name required.') }
      if (!this.companyCopy.address) { this.errors.push('Address required.') }
      if (!this.companyCopy.zipCode) { this.errors.push('ZIP required.') }
      if (!this.companyCopy.city) { this.errors.push('City required.') }
      if (!this.companyCopy.country) { this.errors.push('Country required.') }

      if (this.errors.length > 0) { return }

      const payload = this.companyCopy

      try {
        let response
        if (this.isNew) {
          response = await this.$axios.$post('/companies', payload)

          this.$router.replace({ path: '/companies/' + response.company.id })

          this.$toast.open({
            message: 'Successfully created',
            position: 'top-right'
          })

          return
        } else {
          response = await this.$axios.$patch('/companies/' + this.$route.params.id, payload)
        }

        this.company = Object.assign({}, response.company)
        this.$toast.open({
          message: 'Saved successfully',
          position: 'top-right'
        })
        this.isEdit = false
      } catch (e) {
        this.$toast.open({
          message: e,
          position: 'top-right',
          type: 'error'
        })
      }
    },
    async deleteClicked () {
      if (confirm('Do you really want to delete company?')) {
        try {
          await this.$axios.$delete('/companies/' + this.$route.params.id)

          this.$router.replace({ path: '/' })

          this.$toast.open({
            message: 'Deleted successfully',
            position: 'top-right'
          })
        } catch (e) {
          this.$toast.open({
            message: e,
            position: 'top-right',
            type: 'error'
          })
        }
      }
    },
    async addOwner (e) {
      e.preventDefault()

      this.errors = []

      if (!this.newOwner.name) { this.errors.push('Owner name is required.') }
      if (!this.newOwner.title) { this.errors.push('Title name is required.') }

      if (this.errors.length) { return }

      try {
        const payload = this.newOwner
        await this.$axios.$post('/companies/' + this.$route.params.id + '/owners', payload)
        this.newOwner = {}

        this.getCompany()

        this.$toast.open({
          message: 'Added successfully',
          position: 'top-right'
        })
      } catch (e) {
        this.$toast.open({
          message: e,
          position: 'top-right',
          type: 'error'
        })
      }
    },
    async removeOwner (owner) {
      const msg = 'Do you really want to remove ' + owner.name + ' as owner?'
      if (confirm(msg)) {
        try {
          await this.$axios.$delete('/companies/' + this.$route.params.id + '/owners/' + owner.id)

          this.getCompany()

          this.$toast.open({
            message: 'Removed successfully',
            position: 'top-right'
          })
        } catch (e) {
          this.$toast.open({
            message: e,
            position: 'top-right',
            type: 'error'
          })
        }
      }
    }
  },
  head () {
    return {
      title: 'Cool company service'
    }
  }
}
</script>
