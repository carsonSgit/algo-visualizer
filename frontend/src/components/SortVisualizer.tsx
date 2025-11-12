'use client';
import React, { useState, useEffect } from 'react';
import { Button } from './ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './ui/card';
import { Slider } from './ui/slider';
import { Play, Pause, RotateCcw, Shuffle } from 'lucide-react';

interface Step {
  array: number[];
  comparing: number[];
  swapped: number[];
  sorted: number[];
  step_number: number;
  message: string;
}

interface SortResult {
  steps: Step[];
  algorithm: string;
  duration: string;
  comparisons: number;
  swaps: number;
}

export const SortVisualizer: React.FC = () => {
  const [array, setArray] = useState<number[]>([]);
  const [steps, setSteps] = useState<Step[]>([]);
  const [currentStep, setCurrentStep] = useState(0);
  const [isPlaying, setIsPlaying] = useState(false);
  const [speed, setSpeed] = useState(200);
  const [isLoading, setIsLoading] = useState(false);
  const [arraySize, setArraySize] = useState(20);

  useEffect(() => {
    generateArray();
  }, [arraySize]);

  useEffect(() => {
    if (!isPlaying) return;

    if (currentStep >= steps.length - 1) {
      setIsPlaying(false);
      return;
    }

    const timeout = setTimeout(() => {
      setCurrentStep(prev => prev + 1);
    }, 300 - speed);

    return () => clearTimeout(timeout);
  }, [isPlaying, currentStep, steps, speed]);

  const apiBase = process.env.NEXT_PUBLIC_API_BASE ?? '';

  const generateArray = async () => {
    try {
      setIsLoading(true);
      const response = await fetch(`${apiBase}/api/generate?size=${arraySize}`);
      const data = await response.json();
      setArray(data.array);
      setSteps([]);
      setCurrentStep(0);
      setIsPlaying(false);
    } catch (error) {
      console.error('Error generating array:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const startSort = async () => {
    if (array.length === 0) {
      await generateArray();
      return;
    }

    try {
      setIsLoading(true);
      const response = await fetch(`${apiBase}/api/sort`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ array, algorithm: 'bubble' }),
      });
      const result: SortResult = await response.json();
      setSteps(result.steps);
      setCurrentStep(0);
      setIsPlaying(true);
    } catch (error) {
      console.error('Error starting sort:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const currentStepData = steps[currentStep] || {
    array: array,
    comparing: [],
    swapped: [],
    sorted: [],
    step_number: 0,
    message: 'Ready to sort',
  };

  const maxValue = Math.max(...array, 100);

  return (
    <div className="min-h-screen w-full bg-background">
      <div className="max-w-6xl mx-auto p-6 md:p-8">
        <div className="flex items-end justify-between mb-6">
          <div>
            <h1 className="text-2xl md:text-3xl font-semibold tracking-tight">Algorithm Visualizer</h1>
            <p className="text-sm text-muted-foreground mt-1">Minimal, monochrome sorting visualization</p>
          </div>
          <div className="hidden md:flex items-center gap-3">
            <Button variant="outline" onClick={generateArray} disabled={isLoading || isPlaying} className="h-9 px-3">
              Shuffle
            </Button>
            <Button onClick={startSort} disabled={isLoading || isPlaying} className="h-9 px-3">
              Start
            </Button>
          </div>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-6">
          {/* Visualization */}
          <Card className="lg:col-span-2 bg-card border-border">
            <CardHeader>
              <CardTitle>Bubble Sort</CardTitle>
              <CardDescription>{currentStepData.message}</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="flex items-end justify-center gap-[2px] h-80 bg-muted rounded-md p-4 overflow-hidden">
                {currentStepData.array.map((value, idx) => {
                  const isComparing = currentStepData.comparing.includes(idx);
                  const isSwapped = currentStepData.swapped.includes(idx);
                  const isSorted = currentStepData.sorted.includes(idx);
                  const height = (value / maxValue) * 100;

                  let bgColor = 'bg-zinc-700';
                  if (isSorted) bgColor = 'bg-zinc-300';
                  else if (isSwapped) bgColor = 'bg-zinc-400';
                  else if (isComparing) bgColor = 'bg-zinc-500';

                  return (
                    <div
                      key={idx}
                      className={`flex-1 ${bgColor} transition-all duration-100 rounded-t flex items-end justify-center text-[10px] text-foreground`}
                      style={{ height: `${height}%` }}
                      title={`${value}`}
                    >
                      {currentStepData.array.length <= 20 && <span className="pb-1 opacity-70">{value}</span>}
                    </div>
                  );
                })}
              </div>
            </CardContent>
          </Card>

          {/* Controls */}
          <Card className="bg-card border-border h-fit">
            <CardHeader>
              <CardTitle>Controls</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <label className="text-xs text-muted-foreground uppercase tracking-wider">Array Size</label>
                <Slider
                  value={[arraySize]}
                  onValueChange={(v) => setArraySize(v[0])}
                  min={5}
                  max={100}
                  step={1}
                  disabled={isPlaying || steps.length > 0}
                  className="cursor-pointer"
                />
                <p className="text-xs text-muted-foreground">{arraySize} elements</p>
              </div>

              <div className="space-y-2">
                <label className="text-xs text-muted-foreground uppercase tracking-wider">Speed</label>
                <Slider
                  value={[speed]}
                  onValueChange={(v) => setSpeed(v[0])}
                  min={50}
                  max={290}
                  step={10}
                  className="cursor-pointer"
                />
                <p className="text-xs text-muted-foreground">
                  {speed === 50 ? 'Slow' : speed === 150 ? 'Normal' : speed === 290 ? 'Fast' : 'Custom'}
                </p>
              </div>

              <div className="space-y-2 pt-4">
                <Button
                  onClick={startSort}
                  disabled={isLoading || isPlaying}
                  className="w-full gap-2 h-9"
                >
                  <Play className="w-4 h-4 opacity-80" />
                  Start Sorting
                </Button>

                <Button
                  variant="outline"
                  onClick={() => setIsPlaying(!isPlaying)}
                  disabled={steps.length === 0 || currentStep >= steps.length - 1}
                  className="w-full gap-2 h-9"
                >
                  {isPlaying ? (
                    <>
                      <Pause className="w-4 h-4 opacity-80" />
                      Pause
                    </>
                  ) : (
                    <>
                      <Play className="w-4 h-4 opacity-80" />
                      Resume
                    </>
                  )}
                </Button>

                <Button
                  variant="outline"
                  onClick={() => {
                    setIsPlaying(false);
                    setCurrentStep(0);
                    setSteps([]);
                  }}
                  disabled={steps.length === 0}
                  className="w-full gap-2 h-9"
                >
                  <RotateCcw className="w-4 h-4 opacity-80" />
                  Reset
                </Button>

                <Button
                  variant="outline"
                  onClick={generateArray}
                  disabled={isLoading || isPlaying}
                  className="w-full gap-2 h-9"
                >
                  <Shuffle className="w-4 h-4 opacity-80" />
                  Shuffle Array
                </Button>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Stats */}
        {steps.length > 0 && (
          <Card className="bg-card border-border">
            <CardHeader>
              <CardTitle>Statistics</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-2 md:grid-cols-5 gap-4">
                <div>
                  <p className="text-muted-foreground text-sm">Step</p>
                  <p className="text-2xl font-bold">
                    {currentStep + 1} / {steps.length}
                  </p>
                </div>
                <div>
                  <p className="text-muted-foreground text-sm">Progress</p>
                  <p className="text-2xl font-bold">
                    {Math.round((currentStep / steps.length) * 100)}%
                  </p>
                </div>
                <div>
                  <p className="text-muted-foreground text-sm">Comparisons</p>
                  <p className="text-2xl font-bold">{steps[steps.length - 1]?.step_number || 0}</p>
                </div>
                <div>
                  <p className="text-muted-foreground text-sm">Array Size</p>
                  <p className="text-2xl font-bold">{array.length}</p>
                </div>
                <div>
                  <p className="text-muted-foreground text-sm">Status</p>
                  <p className="text-2xl font-bold">
                    {currentStep >= steps.length - 1 ? 'Complete' : isPlaying ? 'Running' : 'Paused'}
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>
        )}

        {/* Step Slider */}
        {steps.length > 0 && (
          <Card className="mt-6 bg-card border-border">
            <CardContent className="pt-6">
              <div className="space-y-2">
                <label className="text-xs text-muted-foreground uppercase tracking-wider">Step Control</label>
                <Slider
                  value={[currentStep]}
                  onValueChange={(v) => {
                    setIsPlaying(false);
                    setCurrentStep(v[0]);
                  }}
                  min={0}
                  max={steps.length - 1}
                  step={1}
                  className="cursor-pointer"
                />
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  );
};

