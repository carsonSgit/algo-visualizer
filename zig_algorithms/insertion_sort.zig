const std = @import("std");
const io_utils = @import("io_utils.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const parsed = try io_utils.readInput(allocator);
    defer parsed.deinit();
    
    var arr = try allocator.alloc(i64, parsed.value.array.len);
    defer allocator.free(arr);
    @memcpy(arr, parsed.value.array);

    var step_num: usize = 0;
    // For insertion sort, the "sorted" portion is relative (0..i), but we can track specific verified indices if we want.
    // Usually visualization just shows the sorted partition growing.
    // We'll track indices 0..i as "sorted" partition.
    var sorted_indices = std.ArrayList(usize).init(allocator);
    defer sorted_indices.deinit();

    try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, "Starting Insertion Sort");
    step_num += 1;

    // First element is trivially sorted
    try sorted_indices.append(0);

    var i: usize = 1;
    while (i < arr.len) : (i += 1) {
        const key = arr[i];
        var j_signed: isize = @as(isize, @intCast(i)) - 1;
        
        var msg_buf: [100]u8 = undefined;
        const start_msg = try std.fmt.bufPrint(&msg_buf, "Selecting {d} at index {d}", .{key, i});
        try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, start_msg);
        step_num += 1;

        while (j_signed >= 0) {
            const j = @as(usize, @intCast(j_signed));
            
            // Compare
            const compare_indices = [_]usize{j, j + 1};
            const compare_msg = try std.fmt.bufPrint(&msg_buf, "Comparing {d} with {d}", .{arr[j], key});
            try io_utils.printStep(arr, &compare_indices, &.{}, sorted_indices.items, step_num, compare_msg);
            step_num += 1;

            if (arr[j] > key) {
                // Shift
                arr[j + 1] = arr[j];
                const shift_msg = try std.fmt.bufPrint(&msg_buf, "Moving {d} forward", .{arr[j]});
                try io_utils.printStep(arr, &.{}, &compare_indices, sorted_indices.items, step_num, shift_msg);
                step_num += 1;
                j_signed -= 1;
            } else {
                break;
            }
        }
        
        const insert_pos = @as(usize, @intCast(j_signed + 1));
        arr[insert_pos] = key;
        
        try sorted_indices.append(i); // Partition grows
        const place_msg = try std.fmt.bufPrint(&msg_buf, "Inserted {d} at position {d}", .{key, insert_pos});
        try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, place_msg);
        step_num += 1;
    }

    try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, "Sorting complete!");
}
