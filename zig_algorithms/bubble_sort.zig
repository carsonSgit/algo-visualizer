const std = @import("std");
const io_utils = @import("io_utils.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Read input
    const parsed = try io_utils.readInput(allocator);
    defer parsed.deinit();
    
    // We need a mutable copy of the array
    var arr = try allocator.alloc(i64, parsed.value.array.len);
    defer allocator.free(arr);
    @memcpy(arr, parsed.value.array);

    var step_num: usize = 0;
    var sorted_indices = std.ArrayList(usize).init(allocator);
    defer sorted_indices.deinit();

    // Initial State
    try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, "Starting Bubble Sort");
    step_num += 1;

    const n = arr.len;
    var i: usize = 0;
    while (i < n - 1) : (i += 1) {
        var swapped = false;
        var j: usize = 0;
        
        while (j < n - i - 1) : (j += 1) {
            // Compare Step
            const compare_indices = [_]usize{j, j + 1};
            var msg_buf: [64]u8 = undefined;
            const compare_msg = try std.fmt.bufPrint(&msg_buf, "Comparing {d} and {d}", .{arr[j], arr[j+1]});
            
            try io_utils.printStep(arr, &compare_indices, &.{}, sorted_indices.items, step_num, compare_msg);
            step_num += 1;

            if (arr[j] > arr[j + 1]) {
                const temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
                swapped = true;

                // Swap Step
                const swap_msg = try std.fmt.bufPrint(&msg_buf, "Swapped positions {d} and {d}", .{j, j+1});
                try io_utils.printStep(arr, &.{}, &compare_indices, sorted_indices.items, step_num, swap_msg);
                step_num += 1;
            }
        }

        // Mark sorted
        try sorted_indices.append(n - i - 1);

        if (!swapped) {
            // If no swaps, everything remaining is sorted
            var k: usize = 0;
            while (k < n - i - 1) : (k += 1) {
                try sorted_indices.append(k);
            }
            try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, "Array is sorted early!");
            return;
        }
    }
    
    // Add final remaining element (index 0) to sorted
    try sorted_indices.append(0);
    try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, "Sorting complete!");
}
