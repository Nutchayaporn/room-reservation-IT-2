<template>
  <v-hover v-slot="{ hover }">
    <v-card class="rounded-lg" :elevation="hover ? 4 : 0" outlined>
      <v-row no-gutters align="center">
        <v-col md="4" class="py-6 text-3xl text-center">
          <span style="color: #e91e63"> {{ item.room }} </span>
        </v-col>

        <v-col md="8" class="py-6">
          <span> Reserve Date : </span>
          {{ item.reserveDate }}<br />

          <span> Attendee : </span>
          {{ item.attendee }}<br />

          <span> Status : </span>
          <v-chip
            label
            outlined
            class=""
            :color="
              item.status == 'Approved'
                ? '#4CAF50'
                : item.status == 'Pending'
                ? '#FFC107'
                : '#F44336'
            "
          >
            {{ item.status }} </v-chip
          ><br />
        </v-col>
        <v-row v-if="approve" justify="center">
          <v-btn
            color="warning"
            class="px-10"
            rounded
            dark
            @click="readMore(item)"
          >
            More
          </v-btn>
        </v-row>
      </v-row>
    </v-card>
  </v-hover>
</template>

<script>
export default {
  props: ["item", "readMore", "staff", "approve"],
};
</script>

<style></style>
