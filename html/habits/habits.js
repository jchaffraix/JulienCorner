class Day {
  constructor(year, month, day) {
      this.year = year;
      // TODO: Validate that month is in 1..12 (not in Date format but ISO)
      this.month = month;
      this.day = day;
  }

  static today() {
    const d = new Date();
    return new Day(d.getUTCFullYear(), d.getUTCMonth() + 1, d.getUTCDate());
  }

  static ago(days) {
      return Day.today().before(days);
  }

  before(days) {
    let d = new Date(this.toString() + 'T00:00:00Z');
    d.setUTCDate(d.getUTCDate() - days);
    return new Day(d.getUTCFullYear(), d.getUTCMonth() + 1, d.getUTCDate());
  }

  after(days) {
    let d = new Date(this.toString() + 'T00:00:00Z');
    d.setUTCDate(d.getUTCDate() + days);
    return new Day(d.getUTCFullYear(), d.getUTCMonth() + 1, d.getUTCDate());
  }

  toString() {
    // We need to pad to be a valid ISO string.
    let pad = (i) => {
      if (i < 10) {
        return '0' + i;
      }
      return i;
    }
    return this.year + '-' + pad(this.month) + '-' + pad(this.day);
  }

  static fromISOString(isoStr) {
    const parts = isoStr.split('-');
    if (parts.length != 3) {
      throw 'Invalid ISO string: ' + isoStr;
    }

    const parseOrThrow = (str) => {
      const i = parseInt(str)
      if (isNaN(i)) {
        throw "Invalid date: " + isoStr + ", " + str + " is not an int";
      }
      return i
    };
    return new Day(
      parseOrThrow(parts[0]),
      parseOrThrow(parts[1]),
      parseOrThrow(parts[2]));
  }

  isAfterOrSame(other) {
    return this.year >= other.year && this.month >= other.month && this.day >= other.day;
  }

  isBefore(other) {
    return !this.isAfterOrSame(other);
  }
}

class Run {
  constructor(start, length) {
    this.start = start;
    this.length = length;
  }

  includesToday() {
    if (this.length === 0) {
      // A zero-length run is special and never matches today.
      return false;
    }
    return this.start.after(this.length).isAfterOrSame(Day.today());
  }

  isBroken() {
    // It is broken if the run was not done yesterday.
    return this.start.after(this.length).isBefore(Day.ago(1));
  }
}

class HabitRun {
  constructor(name, created_at, start, length) {
    this.name = name;
    this.created_at = created_at;
    this.run = new Run(start, length);
  }

  toJSON() {
    return {
      'name': this.name,
      'created_at': this.created_at.toString(),
      'run_started_at': this.run.start.toString(),
      'run_length': this.run.length,
    };
  }

  static fromJSONObject(json) {
    // Invalid format.
    if (json.name == undefined
      || json.created_at == undefined
      || json.run_started_at == undefined
      || json.run_length == undefined) {
      return null;
    }

    return new HabitRun(
      json.name,
      Day.fromISOString(json.created_at),
      Day.fromISOString(json.run_started_at),
      json.run_length);
   }

   didToday() {
     if (this.run.isBroken()) {
       // Restart the run.
       this.run = new Run(Day.today(), 1);
       return;
     }
     this.run.length++;
   }
}

let habits = [];
const kLocalStorageKey = 'habits';

function saveHabits() {
  console.log("Saving habits: " + JSON.stringify(habits));
  localStorage.setItem(kLocalStorageKey, JSON.stringify(habits));
}

function loadHabits() {
  jsonHabits = JSON.parse(localStorage.getItem(kLocalStorageKey)) || [];
  jsonHabits.forEach(jsonHabit => {
    let habit = HabitRun.fromJSONObject(jsonHabit);
    if (habit === null) {
      // TODO: Stop dropping habits?
      console.log("Dropped invalid habit: " + jsonHabit);
      return;
    }
    habits.push(habit);
  });
}

function didHabitToday(elem, habit) {
  // TODO: Debounce.
  console.log("ticked: " + habit.name);
  habit.didToday();
  saveHabits();
  renderHabits();
}

function renderHabits() {
  let fragment = document.createDocumentFragment();
  habits.forEach(habit => {
    const span = document.createElement('div');
    span.appendChild(document.createTextNode(habit.name));
    span.appendChild(document.createTextNode('('));
    // TODO: DST wrong?
    // TODO: This is a bit too crude.
    if (habit.run.length > 0) {
      if (habit.run.isBroken()) {
        span.appendChild(document.createTextNode('\u274C'));
      } else {
          span.appendChild(document.createTextNode(habit.run.length));
          span.appendChild(document.createTextNode('\u2705'));
      }
    }
    if (!habit.run.includesToday()) {
      const input = document.createElement('input');
      input.type = 'checkbox';
      input.addEventListener('click', (e) => didHabitToday(e, habit));
      span.appendChild(input);
    }
    span.appendChild(document.createTextNode(')'));
    fragment.appendChild(span);
  });

  if (fragment.childNodes.length) {
    // Remove previous habits.
    const habitElem = document.getElementById('habits');
    while ((elem = habitElem.firstElementChild) != null) {
      habitElem.removeChild(elem);
    }
    habitElem.appendChild(fragment);
  }
}

function addHabit() {
  // TODO: Debounce!!!!

  const habit_name = document.getElementById("habit_name_in").value;
  if (habit_name === "") {
    // TODO: Send some validation error.
    return;
  }
  habits.push(new HabitRun(
    habit_name,
    Day.today(),
    Day.today(),
    0,
  ));

  allowAddingHabits();
  saveHabits();
  renderHabits();
}

function loadHandler() {
  document.getElementById('add_habit').addEventListener('click', addHabit);
  loadHabits();
  renderHabits();
  allowAddingHabits();
}

function allowAddingHabits() {
  let disabled = false;
  habits.forEach((habit) => {
    if (habit.run.length < 30) {
      disabled = true;
    }
  });
  if (!disabled) {
    document.getElementById('add_habit_wrapper').classList.add('visible');
  } else {
    document.getElementById('add_habit_wrapper').classList.remove('visible');
  }
}

window.addEventListener('load', loadHandler);
