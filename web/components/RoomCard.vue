<template>
  <v-hover v-slot="{ hover }">
    <v-card :elevation="hover ? 2 : 0" class="p-3 rounded-lg" outlined>
      <v-list-item three-line style="align-items: flex-start">
        <v-list-item-content>
          <v-list-item-title class="mb-1" style="font-size: 1.5rem">
            {{ room.room_name }}
          </v-list-item-title>
          <v-list-item-subtitle
            >Type : {{ room.room_type.type_name }}</v-list-item-subtitle
          >
          <v-list-item-subtitle
            >Capacity : {{ room.room_capacity }}</v-list-item-subtitle
          >
        </v-list-item-content>
        <v-chip
          label
          outlined
          class="mt-2"
          :color="
            room.room_status == 'Available'
              ? 'green'
              : room.room_status == 'Not available'
              ? 'red'
              : 'primary'
          "
        >
          {{ room.room_status }}
        </v-chip>
      </v-list-item>

      <v-card-actions style="justify-content: end">
        <v-btn color="pink" dark @click.stop="selected(room)">
          Book
          <v-icon right>mdi-calendar-multiple-check</v-icon>
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-hover>
</template>

<script>
export default {
  props: ["room", "selected"],
};
</script>

<style></style>
