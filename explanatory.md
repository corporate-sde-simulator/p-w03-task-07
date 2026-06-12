# Beginner Explanatory Guide: PLATFORM-2887: Build windowed data aggregation worker

> **Task Type**: Product Task  
> **Domain/Focus**: Backend Development, Golang, Streaming Data Processing

---

## 1. The Goal (In-Depth Beginner Explanation)

### The Core Problem
The task at hand is to implement a window manager for a real-time data aggregation system. Currently, the system lacks the ability to group incoming data into defined time intervals, known as "windows," which is essential for computing various aggregates such as sums, averages, counts, minimums, and maximums. Without this functionality, the system cannot efficiently process and analyze streaming data, leading to potential delays in data insights and decision-making for users relying on real-time analytics.

The absence of a window manager means that incoming data events are processed individually without any context of time, making it difficult to derive meaningful insights from the data. For example, if a user wants to know the average temperature over the last 10 minutes, the system must be able to collect and aggregate data points that fall within that specific time frame. Implementing the window manager is crucial for enhancing the system's capabilities, ensuring that users can access timely and relevant data insights.

### Jargon Buster (Key Terms Explained)
* **Window Manager**: A component that organizes incoming data into time-based groups (windows) for processing. For instance, if data is collected every second, a window manager can group this data into 1-minute intervals to calculate aggregates like the average value over that minute.
  
* **Aggregation**: The process of summarizing data points into a single value. For example, if you have temperature readings of 20°C, 22°C, and 21°C over a minute, the aggregation (average) would be 21°C. This helps in simplifying large datasets into understandable metrics.

* **Sliding Window**: A type of window that moves forward in time, continuously updating the data it includes. For example, if you have a sliding window of 5 minutes, it will always include the last 5 minutes of data, updating as new data comes in.

* **Tumbling Window**: A fixed-size window that does not overlap. For example, if you have a tumbling window of 10 minutes, it will aggregate data from 0-10 minutes, then from 10-20 minutes, and so on, without any overlap between the windows.

### Expected Outcome
After implementing the solution, the system should be able to effectively manage time windows for incoming data. 

**Before**: The system processes each data event individually, lacking the ability to compute aggregates over time periods. Users cannot derive meaningful insights from the data.

**After**: The system groups incoming data into defined time windows, allowing for the computation of aggregates. Users can now query for averages, sums, and other metrics over specified time intervals, leading to improved data analysis and decision-making.

---

## 2. Related Coding Concepts & Syntax (50% Theory, 50% Practice)

### Concept 1: Data Structures
#### 📘 Theoretical Overview (50%)
* **Why it exists**: Data structures are essential for organizing and storing data efficiently. They allow for quick access and modification of data, which is crucial in a real-time system where data is constantly being received and processed. Without appropriate data structures, the system could become slow and inefficient, leading to delays in data processing.

* **Key Mechanisms**: Common data structures include arrays, lists, maps, and queues. Each has its own strengths and weaknesses. For example, arrays allow for fast access to elements but have a fixed size, while lists can grow dynamically but may have slower access times.

#### 💻 Syntax & Practical Examples (50%)
* **Language Syntax** (Golang):
  ```go
  type Window struct {
      StartTime time.Time
      EndTime   time.Time
      Events    []Event
  }
  ```
  In this example, we define a `Window` struct that holds the start and end times of the window, as well as a slice of `Event` objects that belong to that window.

* **Real-World Application**:
  ```go
  func NewWindow(start time.Time, duration time.Duration) Window {
      return Window{
          StartTime: start,
          EndTime:   start.Add(duration),
          Events:    []Event{},
      }
  }
  ```
  This function creates a new window with a specified start time and duration, initializing an empty list of events.

---

## 3. Step-by-Step Logic & Walkthrough

1. **Step 1: Locate and Analyze the Target File**
   * Navigate to the `p-w03-task-07` folder and open the `WindowManager.go` file. This file contains the `WindowManager` struct and its methods.
   * Focus on the methods marked with TODO comments, specifically `AddEvent()`, `CloseWindow()`, and `GetActiveWindows()`.

2. **Step 2: Input Verification & Validation**
   * Before processing events, ensure that the input data is valid. Check if the event is not null and contains the necessary attributes (e.g., timestamp).

3. **Step 3: Core Implementation / Modification**
   * Implement the `AddEvent()` method to assign incoming events to the correct time window based on their timestamps. If the event falls outside the current window, create a new window.
   * Implement the `CloseWindow()` method to compute aggregates for completed windows. Use the existing `Aggregator` to calculate the required metrics.
   * Implement the `GetActiveWindows()` method to return a list of currently open windows.

4. **Step 4: Output Verification & Testing**
   * Run the unit tests provided in the repository to ensure that all functionalities work as expected. Verify that the tests for `AddEvent()`, `CloseWindow()`, and `GetActiveWindows()` pass successfully.

---

## 4. Detailed Walkthrough of Test Cases

### Test Case 1: Standard / Success Case
* **Description**: This test checks if events are correctly assigned to the appropriate time window.
* **Inputs**:
  ```json
  {
      "events": [
          {"timestamp": "2023-10-01T10:00:00Z", "value": 10},
          {"timestamp": "2023-10-01T10:00:30Z", "value": 20}
      ]
  }
  ```
* **Step-by-Step Execution Trace**:
  1. The `AddEvent()` function receives the first event with a timestamp of 10:00:00.
  2. It checks if there is an existing window for this timestamp. Since there isn't, it creates a new window starting at 10:00:00.
  3. The event is added to the window's event list.
  4. The second event with a timestamp of 10:00:30 is processed similarly and added to the same window.
* **Expected Output**: The window should contain both events, and the aggregates can be computed correctly.

### Test Case 2: Edge Case / Validation Fail
* **Description**: This test checks how the system handles an event with a null timestamp.
* **Inputs**:
  ```json
  {
      "events": [
          {"timestamp": null, "value": 15}
      ]
  }
  ```
* **Step-by-Step Execution Trace**:
  1. The `AddEvent()` function receives the event with a null timestamp.
  2. The validation block detects that the timestamp is invalid.
  3. The function halts execution early and returns an error message indicating the invalid input.
* **Expected Output**: An error message stating "Invalid event timestamp" should be returned, and no events should be added to any window.