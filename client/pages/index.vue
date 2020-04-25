<template>
  <section class="section">
    <div class="container">
      <h1 class="title">
        Companies
      </h1>
      <h2 class="subtitle">
        A list of all our companies
      </h2>

      <nuxt-link :to="'/companies/0'">
        <button class="button is-primary">
          Add company
        </button>
      </nuxt-link>

      <br>

      <br>
      <div class="table-container">
        <table class="table is-fullwidth is-hoverable is-striped">
          <tr>
            <th>Name</th>
            <th>Company ID</th>
            <th>Address</th>
            <th>City</th>
            <th>ZIP</th>
            <th>Country</th>
            <th>E-mail</th>
            <th>Phone</th>
          </tr>
          <tr v-for="(company, index) in companies" :key="index" class="company-row" @click="() => onCompanyClick(company.id)">
            <td>{{ company.name }}</td>
            <td>{{ company.companyID }}</td>
            <td>{{ company.address }}</td>
            <td>{{ company.city }}</td>
            <td>{{ company.zipCode }}</td>
            <td>{{ company.country }}</td>
            <td>{{ company.email }}</td>
            <td>{{ company.phone }}</td>
          </tr>
          <tr v-if="!companies.length">
            <td colspan="9">
              <i>No companies</i>
            </td>
          </tr>
        </table>
      </div>
    </div>
  </section>
</template>

<script>

export default {
  data () {
    return {
      companies: []
    }
  },
  created () {
    this.getCompanies()
  },
  methods: {
    async getCompanies () {
      const data = await this.$axios.$get('/companies')

      console.log('data', data)

      this.companies = data.companies
    },
    onCompanyClick (id) {
      this.$router.push('companies/' + id)
    }
  },
  head () {
    return {
      title: 'Cool company service'
    }
  }
}
</script>

<style lang="scss" scoped>
  .company-row {
    cursor: pointer;
  }
</style>
