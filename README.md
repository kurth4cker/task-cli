# task-cli-go - Manage your tasks [WIP]

**task-cli-go** is a simple CLI program. It is inspired by [roadmap.sh][].

**Note:** This program is still in Work-in-Progress. Most of functionality is
not implemented yet.

**task-cli-go** creates task database as a JSON file in running directory called
`tasks.json`.

## Usage

Here are some examples:

```sh
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Copying

**task-cli-go** is licensed under MPL-2.0. See file COPYING for details.
