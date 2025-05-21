<template>
  <div class="calendar-container">
    <div class="calendar-header">
      <button @click="prevMonth">&lt;</button>
      <h2>{{ monthYear }}</h2>
      <button @click="nextMonth">&gt;</button>
    </div>
    <div class="calendar-grid">
      <div class="calendar-day-name" v-for="day in weekDays" :key="day">{{ day }}</div>
      <div
        v-for="day in calendarDays"
        :key="day.date"
        :class="['calendar-day', { 'calendar-day--disabled': !day.inCurrentMonth, 'calendar-day--booked': isBooked(day.date) }]"
      >
        {{ day.day }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    bookings: {
      type: Array,
      required: true,
      
    },
  },
  data() {
    return {
      currentDate: new Date(),
      weekDays: ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс'],
    };
  },
  computed: {
    monthYear() {
      return this.currentDate.toLocaleString('ru-RU', { year: 'numeric', month: 'long' });
    },
    calendarDays() {
      const year = this.currentDate.getFullYear();
      const month = this.currentDate.getMonth();

      
      const firstDay = new Date(year, month, 1);
      
      let startDay = firstDay.getDay();
      if (startDay === 0) startDay = 7; // воскресенье в JS = 0, сдвигаем на 7

      // Количество дней в месяце
      const daysInMonth = new Date(year, month + 1, 0).getDate();

      // Количество дней в предыдущем месяце для заполнения пустых ячеек
      const prevMonthDaysCount = startDay - 1;

      // Дни предыдущего месяца
      const prevMonthLastDate = new Date(year, month, 0).getDate();

      // Создаём массив дней календаря (6 недель по 7 дней = 42 дня)
      const totalDays = 42;
      const days = [];

      for (let i = 0; i < totalDays; i++) {
        let dayNumber = i - prevMonthDaysCount + 1;
        let inCurrentMonth = true;
        let date;

        if (dayNumber <= 0) {
          // Дни предыдущего месяца
          inCurrentMonth = false;
          date = new Date(year, month - 1, prevMonthLastDate + dayNumber);
          dayNumber = date.getDate();
        } else if (dayNumber > daysInMonth) {
          // Дни следующего месяца
          inCurrentMonth = false;
          date = new Date(year, month + 1, dayNumber - daysInMonth);
          dayNumber = date.getDate();
        } else {
          // Текущий месяц
          date = new Date(year, month, dayNumber);
        }

        days.push({
          day: dayNumber,
          date: date.toISOString().slice(0, 10), // формат YYYY-MM-DD
          inCurrentMonth,
        });
      }
      return days;
    },
  },
  methods: {
    isBooked(dateStr) {
      return this.bookings.includes(dateStr);
    },
    prevMonth() {
      this.currentDate = new Date(this.currentDate.getFullYear(), this.currentDate.getMonth() - 1, 1);
    },
    nextMonth() {
      this.currentDate = new Date(this.currentDate.getFullYear(), this.currentDate.getMonth() + 1, 1);
    },
  },
};
</script>

<style scoped>
.calendar-container {
  width: 350px;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 10px;
  font-family: Arial, sans-serif;
  user-select: none;
}

.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.calendar-header button {
  background-color: #007bff;
  border: none;
  color: white;
  padding: 5px 12px;
  border-radius: 4px;
  cursor: pointer;
}

.calendar-header h2 {
  margin: 0;
  font-size: 1.2em;
  text-transform: capitalize;
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 5px;
}

.calendar-day-name {
  font-weight: bold;
  text-align: center;
  padding: 5px 0;
  background-color: #f0f0f0;
  border-radius: 4px;
}

.calendar-day {
  height: 40px;
  line-height: 40px;
  text-align: center;
  border-radius: 4px;
  cursor: default;
  transition: background-color 0.3s ease;
}

.calendar-day--disabled {
  color: #bbb;
}

.calendar-day--booked {
  background-color: #e74c3c;
  color: white;
  font-weight: bold;
  cursor: not-allowed;
}
</style>
