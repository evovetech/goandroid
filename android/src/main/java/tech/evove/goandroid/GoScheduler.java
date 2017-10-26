package tech.evove.goandroid;

import java.util.concurrent.TimeUnit;

import core.Core;
import io.reactivex.Scheduler;
import io.reactivex.disposables.Disposable;

public class GoScheduler extends Scheduler {
    private final core.Scheduler scheduler = Core.goScheduler();

    @Override
    public Worker createWorker() {
        return new GoWorker(scheduler.createWorker());
    }

    @Override
    public Disposable scheduleDirect(Runnable run, long delay, TimeUnit unit) {
        return scheduler.schedule(run, unit.toNanos(delay));
    }
}
